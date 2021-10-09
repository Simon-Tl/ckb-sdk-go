package address

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateShortAddress(t *testing.T) {
	shortAddress, _ := GenerateShortAddress(Mainnet)
	fmt.Println(shortAddress)
}

func TestGenerateAcpAddress(t *testing.T) {
	address := "ckt1qyqqtg06h75ymw098r3w0l3u4xklsj04tnsqctqrmc"
	acpAddress, err := GenerateAcpAddress(address)
	assert.Nil(t, err)
	assert.Equal(t, "ckt1qypqtg06h75ymw098r3w0l3u4xklsj04tnsqkm65q6", acpAddress)
}

func TestGenerateChequeAddress(t *testing.T) {
	senderAddress := "ckt1qyq27z6pccncqlaamnh8ttapwn260egnt67ss2cwvz"
	receiverAddress := "ckt1qyqqtg06h75ymw098r3w0l3u4xklsj04tnsqctqrmc"
	acpAddress, err := GenerateChequeAddress(senderAddress, receiverAddress)
	assert.Nil(t, err)
	assert.Equal(t, "ckt1q3sdtuu7lnjqn3v8ew02xkwwlh4dv5x2z28shkwt8p2nfruccux4k5kw5xmckqjq7gwpe990sn88xssv96try4l46hu6nnudr2huau238a4prwus9pqts3uptms", acpAddress)
}

func TestBech32mTypeFullTestnetAddressGenerate(t *testing.T) {
	script := &types.Script{
		CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
		HashType: types.HashTypeType,
		Args:     common.FromHex("0xb39bbc0b3673c7d36450bc14cfcdad2d559c6c64"),
	}

	address, err := GenerateBech32mFullAddress(Mainnet, script)
	println(address)
	if err != nil {
		return
	}

	assert.Equal(t,
		"ckb1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsqdnnw7qkdnnclfkg59uzn8umtfd2kwxceqxwquc4",
		address)

	t.Run("parse full address", func(t *testing.T) {
		mnAddress, err := Parse(address)

		assert.Nil(t, err)
		assert.Equal(t, Mainnet, mnAddress.Mode)
		assert.Equal(t, TypeFull, mnAddress.Type)
		assert.Equal(t, script.CodeHash, mnAddress.Script.CodeHash)
		assert.Equal(t, script.HashType, mnAddress.Script.HashType)
		assert.Equal(t, script.Args, mnAddress.Script.Args)
	})
}
