To generate a wrapper file around the SFConract execute the following commands.

```bash
solc -o $PWD/build --optimize --optimize-runs=2000 --abi --bin-runtime --allow-paths $PWD/contracts --overwrite $PWD/contracts/sfc/Staker.sol

solc -o $PWD/build --optimize --optimize-runs=2000 --bin --bin-runtime --allow-paths $PWD/contracts --overwrite $PWD/contracts/sfc/Staker.sol

TODO: add command from laptop abigen
```

