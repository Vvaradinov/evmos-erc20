KEYRING="test"
HOME_DIR="/mnt/c/evmos/evmos-erc20/node"

# Get the private keys of both users
VLAD_PK=$(evmosd keys unsafe-export-eth-key vlad_validator --home=$HOME_DIR --keyring-backend $KEYRING)
NICK_PK=$(evmosd keys unsafe-export-eth-key nick_validator --home=$HOME_DIR --keyring-backend $KEYRING)
# Get the Bech32 addresses of both users
VLAD_BECH32=$(evmosd keys show vlad_validator  --home=$HOME_DIR --keyring-backend $KEYRING | grep 'address' | grep -o 'evmos[0-9a-z]*')
NICK_BECH32=$(evmosd keys show nick_validator --home=$HOME_DIR --keyring-backend $KEYRING | grep 'address' | grep -o 'evmos[0-9a-z]*')
# Get the Hex addresses of both users
VLAD_HEX=$(evmosd keys parse "$VLAD_BECH32" | grep 'bytes' | grep -o '[0-9A-Z]*')
NICK_HEX=$(evmosd keys parse "$NICK_BECH32" | grep 'bytes' | grep -o '[0-9A-Z]*')

echo Vlad Hex key - $VLAD_HEX
echo Nick Hex key - $NICK_HEX
echo Vlad Private Key - $VLAD_PK
echo Nick Private Key - $NICK_PK

