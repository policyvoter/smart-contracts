rm bindings.go
npx hardhat export-abi
abigen --abi abi/PolicyVoterV0.json --pkg pv0_smartcontract --out bindings.go
rm -rf abi/
echo "generated bindings.go"
