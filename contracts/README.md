To generate a wrapper file around the SFConract execute the following commands.

1. Install solc 0.5.12
2. Execute followings commands.

```bash
solc -o $PWD/build --optimize --optimize-runs=2000 --abi --bin-runtime --allow-paths $PWD/contracts --overwrite $PWD/contracts/sfc/Staker.sol

solc -o $PWD/build --optimize --optimize-runs=2000 --bin --bin-runtime --allow-paths $PWD/contracts --overwrite $PWD/contracts/sfc/Staker.sol

abigen --bin=$PWD/contracts/build/Stakers.bin --abi=$PWD/contracts/build/Stakers.abi --pkg=Stakers --out=contracts/stakers.go
```

