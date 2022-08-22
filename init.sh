# customize the name of your key, the chain-id, moniker of the node, keyring backend, and log level
KEY_VLAD="vlad_validator"
KEY_NICK="nick_validator"
CHAINID="evmos_9000-4"
MONIKER="localtestnet-1"
KEYRING="test"
LOGLEVEL="info"
HOME_DIR="/mnt/c/evmos/evmos-erc20/node"

# Create the initial chain configuration
evmosd init $MONIKER --chain-id=$CHAINID --overwrite --home=$HOME_DIR

# Add my personal key to the keyring
evmosd keys add $KEY_VLAD --keyring-backend test --home=$HOME_DIR

# Add my friend Nikc's key to the keyring
evmosd keys add $KEY_NICK --keyring-backend test --home=$HOME_DIR

# Allocate genesis accounts (cosmos formatted addresses)
evmosd add-genesis-account $KEY_VLAD 100000000000000000000000000stake,100000000000000000000000aphoton --keyring-backend $KEYRING --home=$HOME_DIR

# Sign genesis transaction
evmosd gentx $KEY_VLAD 1000000000000000000000stake --keyring-backend $KEYRING --chain-id $CHAINID --home=$HOME_DIR

# Add the gentx to the genesis file
evmosd collect-gentxs --home=$HOME_DIR --keyring-backend $KEYRING --chain-id $CHAINID

# Check correctness of the genesis.json file
evmosd validate-genesis --home=$HOME_DIR --keyring-backend $KEYRING --chain-id $CHAINID

# Run the local node
evmosd start --home=$HOME_DIR --keyring-backend $KEYRING --chain-id $CHAINID
