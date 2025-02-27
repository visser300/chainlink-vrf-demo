// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "forge-std/Script.sol";
import "../src/RandomNumberConsumer.sol";

contract DeployRandomNumberConsumer is Script {
    function run() external {
        // Get your deployer private key from an environment variable
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        // Sepolia Chainlink VRF configuration (set these values in your .env file)
        uint64 subscriptionId = uint64(vm.envUint("SUBSCRIPTION_ID")); // Your Sepolia subscription id
        // This is the Chainlink VRF Coordinator address for Sepolia
        address vrfCoordinator = 0x8103B0A8A00be2DDC778e6e7eaa21791Cd364625;
        bytes32 keyHash = vm.envBytes32("KEYHASH"); // Your key hash from Chainlink for Sepolia

        vm.startBroadcast(deployerPrivateKey);
        // Deploy the RandomNumberConsumer contract
        RandomNumberConsumer consumer = new RandomNumberConsumer(subscriptionId, vrfCoordinator, keyHash);
        vm.stopBroadcast();
    }
} 