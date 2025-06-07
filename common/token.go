package common

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// GetAllPositionNftAccountByOwner retrieves all position NFT accounts owned by a user
func GetAllPositionNftAccountByOwner(
	ctx context.Context,
	rpcClient *rpc.Client,
	user solana.PublicKey,
) ([]PositionNftAccount, error) {
	// Get all token accounts owned by the user
	tokenAccounts, err := rpcClient.GetTokenAccountsByOwner(
		ctx,
		user,
		&rpc.GetTokenAccountsConfig{
			ProgramId: &solana.Token2022ProgramID,
		},
		&rpc.GetTokenAccountsOpts{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get token accounts: %w", err)
	}

	var userPositionNftAccounts []PositionNftAccount

	// Iterate through token accounts
	for _, tokenAccount := range tokenAccounts.Value {
		data := tokenAccount.Account.Data.GetBinary()
		if len(data) < 165 { // Minimum size for token account data
			continue
		}

		// Parse token account data
		layout := TokenAccountLayout{
			Mint:            solana.PublicKeyFromBytes(data[0:32]),
			Owner:           solana.PublicKeyFromBytes(data[32:64]),
			Amount:          binary.LittleEndian.Uint64(data[64:72]),
			Delegate:        solana.PublicKeyFromBytes(data[72:104]),
			State:           data[104],
			IsNative:        binary.LittleEndian.Uint64(data[105:113]),
			DelegatedAmount: binary.LittleEndian.Uint64(data[113:121]),
			CloseAuthority:  solana.PublicKeyFromBytes(data[121:153]),
		}

		// Check if the account has exactly 1 token (NFT)
		if layout.Amount == 1 {
			userPositionNftAccounts = append(userPositionNftAccounts, PositionNftAccount{
				PositionNft:        layout.Mint,
				PositionNftAccount: tokenAccount.Pubkey,
			})
		}
	}

	return userPositionNftAccounts, nil
}
