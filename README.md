# Meteora DAMM V2 Examples in Go

## Overview

This repository contains examples of how to use the Meteora DAMM V2 program in Go. Powered by [solana-go](https://github.com/gagliardetto/solana-go).

## Prerequisites

- [Download Go](https://go.dev/doc/install)

## Usage

1. Install dependencies

```bash
go mod tidy
```

2. Run the examples

Before running the examples, you need to:

1. Set the private keys and public keys in the examples.
2. Set the RPC endpoint in the examples.
3. Uncomment the `main()` function in the examples.

```bash
go run examples/<file-name>.go
```

## Examples

- [Claim position fee](./examples/claim_position_fee.go)
- [Get all position NFT accounts by owner](./examples/get_all_position_nft_account_by_owner.go)
- [Get pool](./examples/get_pool.go)
- [Get position](./examples/get_position.go)
- [Get positions by user](./examples/get_positions_by_user.go)
- [Get unclaim reward](./examples/get_unclaim_reward.go)
- [Get user position by pool](./examples/get_user_position_by_pool.go)
