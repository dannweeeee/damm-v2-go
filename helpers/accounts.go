package helpers

import (
	"github.com/dannwee/dbc-go/common"
	"github.com/gagliardetto/solana-go"
)

// Derives the event authority PDA
func DeriveEventAuthorityPDA() solana.PublicKey {
	seeds := [][]byte{[]byte("__event_authority")}
	address, _, err := solana.FindProgramAddress(seeds, solana.MustPublicKeyFromBase58(common.DammV2ProgramID))
	if err != nil {
		panic(err)
	}
	return address
}

// Derives the pool authority PDA
func DerivePoolAuthorityPDA() solana.PublicKey {
	seeds := [][]byte{[]byte("pool_authority")}
	address, _, err := solana.FindProgramAddress(seeds, solana.MustPublicKeyFromBase58(common.DammV2ProgramID))
	if err != nil {
		panic(err)
	}
	return address
}
