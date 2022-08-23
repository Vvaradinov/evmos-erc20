package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestObtainClient(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want *ethclient.Client
	}{
		{
			name: "Test with missing url",
			url:  "",
			want: nil,
		},
		{
			name: "Test with a non existing IPC endpoint file path",
			url:  "/home/test/test12345.ipc",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ObtainClient(tt.url); got != tt.want {
				t.Errorf("ObtainClient() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestObtainClientAndTxSigner(t *testing.T) {
	tests := []struct {
		name       string
		rpcUrl     string
		userPK     string
		wantClient *ethclient.Client
		wantAuth   *bind.TransactOpts
	}{
		{
			name:       "Test with missing private key and missing rpc url",
			rpcUrl:     "",
			userPK:     "",
			wantClient: nil,
			wantAuth:   nil,
		}, {
			name:     "Test with missing private key only",
			rpcUrl:   "http://localhost:8545",
			userPK:   "",
			wantAuth: nil,
		}, {
			name:       "Test with missing client only",
			rpcUrl:     "",
			userPK:     "B1FF8C4E3EE8AC07A0C05B08096DCF0E1795BD1A2B6A12F58A9CEAD80206D4CD",
			wantClient: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testAuth, testClient, _ := ObtainClientAndTxSigner(tt.rpcUrl, tt.userPK)
			if testAuth != tt.wantAuth {
				t.Errorf("Test auth = %v, Want auth = %v", testAuth, tt.wantAuth)
			}
			if testClient != tt.wantClient {
				t.Errorf("Test client = %v, Want client = %v", testClient, tt.wantClient)
			}
		})
	}
}
