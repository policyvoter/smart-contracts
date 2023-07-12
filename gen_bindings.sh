# ------ POLICYVOTER ------
rm bindings.go
npx hardhat export-abi
abigen --abi abi/PolicyVoterV1.json --pkg pv1_smartcontract --out bindings.go
rm -rf abi/
echo "generated bindings.go"
echo "remember to change the package name to eth  when you copy paste"


# ------ CHOICES ------

# rm bindings_choices.go
# npx hardhat export-abi
# abigen --abi abi/ChoicesV0.json --pkg choices --out bindings_choices.go
# rm -rf abi/
# echo "generated bindings_choices.go"
# echo "remember to change the package name to `eth` when you copy paste"
