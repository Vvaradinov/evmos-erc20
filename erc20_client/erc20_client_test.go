package erc20_client

import (
	"math/big"
	"reflect"
	"testing"
)

func TestQueryBalance(t *testing.T) {
	tests := []struct {
		name         string
		contractAddr string
		walletAddr   string
		wantBalance  *big.Int
	}{
		{
			name:         "Test with empty strings",
			contractAddr: "",
			walletAddr:   "",
			wantBalance:  nil,
		}, {
			name:         "Test with missing contract address",
			contractAddr: "",
			walletAddr:   "0x6175270C6dfc17C772EEf5170d663342C9482Da7",
			wantBalance:  nil,
		}, {
			name:         "Test with missing wallet address",
			contractAddr: "0xBd7D4f6576c4e14470b0294649d4236a590E2258",
			walletAddr:   "",
			wantBalance:  nil,
		}, {
			name:         "Test with badly formatted contract address",
			contractAddr: "fhki81237asf7y12h1kaosd91368",
			wantBalance:  nil,
		}, {
			name:        "Test with badly formatted wallet address",
			walletAddr:  "hkishd8313y7s7gsd7f1",
			wantBalance: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			balance, _ := QueryBalance(tt.contractAddr, tt.walletAddr)
			if balance != tt.wantBalance {
				t.Errorf("QueryBalance() = %v, want %v", balance, tt.wantBalance)
			}
		})
	}
}

func TestTransfer(t *testing.T) {
	//success := types.ReceiptStatusSuccessful
	//fail := types.ReceiptStatusFailed
	tests := []struct {
		name         string
		contractAddr string
		fromPK       string
		toAddr       string
		amount       string
		wantStatus   *uint64
	}{
		{
			name:         "Test with empty strings",
			contractAddr: "",
			fromPK:       "",
			toAddr:       "",
			amount:       "",
			wantStatus:   nil,
		}, {
			name:         "Test with missing addresses",
			contractAddr: "",
			fromPK:       "",
			toAddr:       "",
			amount:       "100000",
			wantStatus:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, _ := Transfer(tt.contractAddr, tt.fromPK, tt.toAddr, tt.amount)
			if !reflect.DeepEqual(status, tt.wantStatus) {
				t.Errorf("Transfer() = %v, want %v", status, tt.wantStatus)
			}
		})
	}
}
