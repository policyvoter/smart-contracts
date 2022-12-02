// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";

//ChoicesV0 is storing the onboard choices
contract ChoicesV0 is Ownable, Pausable {
    mapping(uint256 => uint256) public choicesTracker;
    mapping(string => bool) public usedKeys;
    string public version = "0.0.1";

    event ChoiceStored(string key, uint256 choice);

    constructor() {}

    /**
		@notice stores the choice in the blockchain
		@param keys - unique key
		@param choices - user choice
	 */
    function storeChoice(string[] calldata keys, uint256[] calldata choices)
        external
        onlyOwner
    {
        require(keys.length == choices.length, "not same length");

        uint256 len = keys.length;
        for (uint256 i = 0; i < len; ) {
            require(!usedKeys[keys[i]], "address used");
            usedKeys[keys[i]] = true;

            choicesTracker[choices[i]] = choicesTracker[choices[i]] + 1;
            emit ChoiceStored(keys[i], choices[i]);

            unchecked {
                ++i;
            }
        }
    }

    //in case this is not needed, the owner can pause this contract
    function setPaused(bool _setPaused) external onlyOwner {
        return (_setPaused) ? _pause() : _unpause();
    }
}
