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

- [Create a pool and swap SOL](./examples/create_pool_and_swap_sol.go)
- [Create a pool and swap USDC](./examples/create_pool_and_swap_usdc.go)
- [Claim creator trading fee](./examples/claim_creator_trading_fee.go)
- [Claim partner trading fee](./examples/claim_partner_trading_fee.go)
- [Fetch pool configuration](./examples/get_pool_config.go)
- [Fetch pool fee metrics](./examples/get_pool_fee_metrics.go)
- [Fetch pool](./examples/get_pool.go)
