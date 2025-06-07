package instructions

import (
	"bytes"
	"context"
	"fmt"

	"github.com/dannwee/dbc-go/common"
	"github.com/dannwee/dbc-go/helpers"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func GetPool(ctx context.Context, poolAddress solana.PublicKey, rpcClient *rpc.Client) (*common.Pool, error) {
	account, err := rpcClient.GetAccountInfo(ctx, poolAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool account: %w", err)
	}

	if account == nil || account.Value == nil {
		return nil, fmt.Errorf("pool account not found")
	}

	data := account.Value.Data.GetBinary()

	if len(data) < 8 {
		return nil, fmt.Errorf("data too short")
	}

	expectedDiscriminator := []byte{241, 154, 109, 4, 17, 177, 109, 188}
	if !bytes.Equal(data[:8], expectedDiscriminator) {
		return nil, fmt.Errorf("invalid discriminator, not a pool account")
	}

	return helpers.DeserializePool(data)
}
