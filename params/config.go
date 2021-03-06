// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// CanonicalDepth is the confirmation required for a block to be considered canonical.
	// Also the distant from the ThangLong checkpoint block to it's snapshot.
	CanonicalDepth = 7
)

// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash = common.HexToHash("0x080eeb525df0e852343ba13afedf2b256f0991c1b18797e18863fd7b4ab3574b")
	TestnetGenesisHash = common.HexToHash("0xe53f0441b5f887499e49cb628262c4acc897492a4fe3236bb3c68023ea725f0d")
	RinkebyGenesisHash = common.HexToHash("0x6341fd3daf94b748c72ced5a5b26028f2474f5f00d824504e4fa37a75767e177")
	BurnAddress        = common.HexToAddress("0x0000000000000000000000000000000000000000")
	TokenAddress       = common.HexToAddress("0x2c783ad80ff980ec75468477e3dd9f86123ecbda") // NTF token contract address
)

var (
	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(66666),
		HomesteadBlock:      big.NewInt(1),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(4),
		ConstantinopleBlock: big.NewInt(15360000),
		PetersburgBlock:     big.NewInt(15360000),
		Dccs: &DccsConfig{
			Period: 2,
			Epoch:  30000,
			// Governance contract
			Contract: common.HexToAddress("0x0000000000000000000000000000000000012345"),
			// ThangLong hard-fork
			StakeRequire:    50000,
			StakeLockHeight: 30000,
			ThangLongBlock:  big.NewInt(15360000),
			ThangLongEpoch:  3000,
		},
	}

	// MainnetTrustedCheckpoint contains the light client trusted checkpoint for the main network.
	MainnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "mainnet",
		SectionIndex: 208,
		SectionHead:  common.HexToHash("0x5e9f7696c397d9df8f3b1abda857753575c6f5cff894e1a3d9e1a2af1bd9d6ac"),
		CHTRoot:      common.HexToHash("0x954a63134f6897f015f026387c59c98c4dae7b336610ff5a143455aac9153e9d"),
		BloomRoot:    common.HexToHash("0x8006c5e44b14d90d7cc9cd5fa1cb48cf53697ee3bbbf4b76fdfa70b0242500a9"),
	}

	// TestnetChainConfig contains the chain parameters to run a node on the Dccs test network.
	TestnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(111111),
		HomesteadBlock:      big.NewInt(1),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(4),
		ConstantinopleBlock: big.NewInt(300000),
		PetersburgBlock:     big.NewInt(300000),
		Dccs: &DccsConfig{
			Period: 2,
			Epoch:  30000,
			// Governance contract
			Contract: common.HexToAddress("0x0000000000000000000000000000000000012345"),
			// ThangLong hard-fork
			StakeRequire:    500,
			StakeLockHeight: 3000,
			ThangLongBlock:  big.NewInt(300000),
			ThangLongEpoch:  300,
		},
	}

	// TestnetTrustedCheckpoint contains the light client trusted checkpoint for the Ropsten test network.
	TestnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "testnet",
		SectionIndex: 139,
		SectionHead:  common.HexToHash("0x9fad89a5e3b993c8339b9cf2cbbeb72cd08774ea6b71b105b3dd880420c618f4"),
		CHTRoot:      common.HexToHash("0xc815833881989c5d2035147e1a79a33d22cbc5313e104ff01e6ab405bd28b317"),
		BloomRoot:    common.HexToHash("0xd94ee9f3c480858f53ec5d059aebdbb2e8d904702f100875ee59ec5f366e841d"),
	}

	// RinkebyChainConfig contains the chain parameters to run a node on the Rinkeby test network.
	RinkebyChainConfig = &ChainConfig{
		ChainID:             big.NewInt(4),
		HomesteadBlock:      big.NewInt(1),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x9b095b36c15eaf13044373aef8ee0bd3a382a5abb92e402afa44b8249c3a90e9"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(1035301),
		ConstantinopleBlock: big.NewInt(3660663),
		PetersburgBlock:     big.NewInt(9999999), //TODO! Insert Rinkeby block number
		Clique: &CliqueConfig{
			Period: 15,
			Epoch:  30000,
		},
	}

	// DccsChainConfig contains the chain parameters to run a node on the Nexty test network.
	DccsChainConfig = &ChainConfig{
		ChainID:             big.NewInt(66666),
		HomesteadBlock:      big.NewInt(1),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x9b095b36c15eaf13044373aef8ee0bd3a382a5abb92e402afa44b8249c3a90e9"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(1035301),
		ConstantinopleBlock: nil,
		PetersburgBlock:     nil,
		Dccs: &DccsConfig{
			Period: 2,
			Epoch:  30000,
		},
	}

	// RinkebyTrustedCheckpoint contains the light client trusted checkpoint for the Rinkeby test network.
	RinkebyTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "rinkeby",
		SectionIndex: 105,
		SectionHead:  common.HexToHash("0xec8147d43f936258aaf1b9b9ec91b0a853abf7109f436a23649be809ea43d507"),
		CHTRoot:      common.HexToHash("0xd92703b444846a3db928e87e450770e5d5cbe193131dc8f7c4cf18b4de925a75"),
		BloomRoot:    common.HexToHash("0xff45a6f807138a2cde0cea0c209d9ce5ad8e43ccaae5a7c41af801bb72a1ef96"),
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Ethash consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, new(EthashConfig), nil, nil}

	// AllCliqueProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Clique consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllCliqueProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, &CliqueConfig{Period: 0, Epoch: 30000}, nil}

	// AllDccsProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Dccs consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllDccsProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, nil, &DccsConfig{Period: 0, Epoch: 30000, ThangLongBlock: big.NewInt(0), ThangLongEpoch: 3000, Contract: common.HexToAddress("0x0")}}

	TestChainConfig = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, new(EthashConfig), nil, nil}

	TestRules = TestChainConfig.Rules(new(big.Int))
)

// TrustedCheckpoint represents a set of post-processed trie roots (CHT and
// BloomTrie) associated with the appropriate section index and head hash. It is
// used to start light syncing from this checkpoint and avoid downloading the
// entire header chain while still being able to securely access old headers/logs.
type TrustedCheckpoint struct {
	Name         string      `json:"-"`
	SectionIndex uint64      `json:"sectionIndex"`
	SectionHead  common.Hash `json:"sectionHead"`
	CHTRoot      common.Hash `json:"chtRoot"`
	BloomRoot    common.Hash `json:"bloomRoot"`
}

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock      *big.Int `json:"byzantiumBlock,omitempty"`      // Byzantium switch block (nil = no fork, 0 = already on byzantium)
	ConstantinopleBlock *big.Int `json:"constantinopleBlock,omitempty"` // Constantinople switch block (nil = no fork, 0 = already activated)
	PetersburgBlock     *big.Int `json:"petersburgBlock,omitempty"`     // Petersburg switch block (nil = same as Constantinople)
	EWASMBlock          *big.Int `json:"ewasmBlock,omitempty"`          // EWASM switch block (nil = no fork, 0 = already activated)

	// Various consensus engines
	Ethash *EthashConfig `json:"ethash,omitempty"`
	Clique *CliqueConfig `json:"clique,omitempty"`
	Dccs   *DccsConfig   `json:"dccs,omitempty"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// DccsConfig is the consensus engine configs for proof-of-foundation based sealing.
// All epoch changing hardforks must have the activate block number divisible by both old and new epoch.
type DccsConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
	// governance smart contract address
	Contract common.Address `json:"contract,omitempty"`
	// Governance contract stake params
	StakeRequire    uint64 `json:"stakeRequire"`    // stake requirement
	StakeLockHeight uint64 `json:"stakeLockHeight"` // lock time (in blocks) after leaving
	// ThangLong hardfork
	ThangLongBlock *big.Int `json:"thangLongBlock,omitempty"` // ThangLong switch block (nil = no fork, 0 = already activated)
	ThangLongEpoch uint64   `json:"thangLongEpoch"`           // Epoch length to reset votes and checkpoint
}

// PositionInEpoch returns the offset of a block from the start of an epoch
func (c *DccsConfig) PositionInEpoch(number uint64) uint64 {
	if c.IsThangLong(new(big.Int).SetUint64(number)) {
		return number % c.ThangLongEpoch
	}
	return number % c.Epoch
}

// IsCheckpoint returns whether a block is at the start of an epoch
func (c *DccsConfig) IsCheckpoint(number uint64) bool {
	return c.PositionInEpoch(number) == 0
}

// Checkpoint returns the epoch block for this block
func (c *DccsConfig) Checkpoint(number uint64) uint64 {
	return number - c.PositionInEpoch(number)
}

// Snapshot returns the snapshot block for this block's epoch.
// Snapshot is (Checkpoint - CanonicalDepth) for ThangLong consensus.
func (c *DccsConfig) Snapshot(number uint64) uint64 {
	// Get the checkpoint first
	cp := c.Checkpoint(number)
	// Get genesis block as checkpoint for 1st epoch
	if cp <= CanonicalDepth {
		return 0
	}
	// Get the state from canonical chain to ensure the chain and state are not in sidefork
	return cp - CanonicalDepth
}

// String implements the stringer interface, returning the consensus engine details.
func (c *DccsConfig) String() string {
	return "dccs"
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var thangLongBlock *big.Int
	var contract string
	var engine interface{}
	switch {
	case c.Ethash != nil:
		engine = c.Ethash
	case c.Clique != nil:
		engine = c.Clique
	case c.Dccs != nil:
		engine = c.Dccs
		thangLongBlock = c.Dccs.ThangLongBlock
		contract = c.Dccs.Contract.Hex()
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v Homestead: %v DAO: %v DAOSupport: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v Constantinople: %v ConstantinopleFix: %v Thang Long: %v Contract: %v Engine: %v}",
		c.ChainID,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		c.ConstantinopleBlock,
		c.PetersburgBlock,
		thangLongBlock,
		contract,
		engine,
	)
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isForked(c.HomesteadBlock, num)
}

// IsDAOFork returns whether num is either equal to the DAO fork block or greater.
func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
	return isForked(c.DAOForkBlock, num)
}

// IsEIP150 returns whether num is either equal to the EIP150 fork block or greater.
func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isForked(c.EIP150Block, num)
}

// IsEIP155 returns whether num is either equal to the EIP155 fork block or greater.
func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isForked(c.EIP155Block, num)
}

// IsEIP158 returns whether num is either equal to the EIP158 fork block or greater.
func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isForked(c.EIP158Block, num)
}

// IsByzantium returns whether num is either equal to the Byzantium fork block or greater.
func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	return isForked(c.ByzantiumBlock, num)
}

// IsConstantinople returns whether num is either equal to the Constantinople fork block or greater.
func (c *ChainConfig) IsConstantinople(num *big.Int) bool {
	return isForked(c.ConstantinopleBlock, num)
}

// IsPetersburg returns whether num is either
// - equal to or greater than the PetersburgBlock fork block,
// - OR is nil, and Constantinople is active
func (c *ChainConfig) IsPetersburg(num *big.Int) bool {
	return isForked(c.PetersburgBlock, num) || c.PetersburgBlock == nil && isForked(c.ConstantinopleBlock, num)
}

// IsEWASM returns whether num represents a block number after the EWASM fork
func (c *ChainConfig) IsEWASM(num *big.Int) bool {
	return isForked(c.EWASMBlock, num)
}

// IsThangLongPreparationBlock returns whether num represents a block number exactly at the ThangLong Preparation
func (c *ChainConfig) IsThangLongPreparationBlock(num *big.Int) bool {
	return c.Dccs != nil && c.Dccs.ThangLongBlock != nil && c.Dccs.IsThangLongPreparationBlock(num)
}

// IsSnapshotBlock returns whether num represents a block number exactly at the snapshot block of an epoch.
func (c *ChainConfig) IsSnapshotBlock(num *big.Int) bool {
	if c.Dccs == nil || c.Dccs.ThangLongBlock == nil {
		return false
	}
	forBlock := new(big.Int).SetUint64(CanonicalDepth)
	forBlock.Add(num, forBlock)
	return c.Dccs.IsThangLong(forBlock) && c.Dccs.IsCheckpoint(forBlock.Uint64())
}

// IsThangLongPreparationBlock returns whether num represents a block number exactly at the ThangLong Preparation
// ThangLong preparation block is hard-coded to 32 blocks before the ThangLong hard-fork
func (c *DccsConfig) IsThangLongPreparationBlock(num *big.Int) bool {
	if c.ThangLongBlock == nil {
		return false
	}
	preparationBlock := big.NewInt(1)
	if c.ThangLongBlock.Cmp(common.Big32) > 0 {
		preparationBlock = preparationBlock.Sub(c.ThangLongBlock, common.Big32)
	}
	return preparationBlock.Cmp(num) == 0
}

// IsThangLong returns whether num represents a block number after the ThangLong fork
func (c *ChainConfig) IsThangLong(num *big.Int) bool {
	return c.Dccs != nil && c.Dccs.IsThangLong(num)
}

// IsThangLong returns whether num represents a block number after the ThangLong fork
func (c *DccsConfig) IsThangLong(num *big.Int) bool {
	return isForked(c.ThangLongBlock, num)
}

// GasTable returns the gas table corresponding to the current phase (homestead or homestead reprice).
//
// The returned GasTable's fields shouldn't, under any circumstances, be changed.
func (c *ChainConfig) GasTable(num *big.Int) GasTable {
	if num == nil {
		return GasTableHomestead
	}
	switch {
	case c.IsThangLong(num):
		return GasTableDccs
	case c.IsConstantinople(num):
		return GasTableConstantinople
	case c.IsEIP158(num):
		return GasTableEIP158
	case c.IsEIP150(num):
		return GasTableEIP150
	default:
		return GasTableHomestead
	}
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if isForkIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, head) {
		return newCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	}
	if isForkIncompatible(c.DAOForkBlock, newcfg.DAOForkBlock, head) {
		return newCompatError("DAO fork block", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if c.IsDAOFork(head) && c.DAOForkSupport != newcfg.DAOForkSupport {
		return newCompatError("DAO fork support flag", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if isForkIncompatible(c.EIP150Block, newcfg.EIP150Block, head) {
		return newCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	}
	if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkIncompatible(c.EIP158Block, newcfg.EIP158Block, head) {
		return newCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	}
	if c.IsEIP158(head) && !configNumEqual(c.ChainID, newcfg.ChainID) {
		return newCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	}
	if isForkIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, head) {
		return newCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	}
	if isForkIncompatible(c.ConstantinopleBlock, newcfg.ConstantinopleBlock, head) {
		return newCompatError("Constantinople fork block", c.ConstantinopleBlock, newcfg.ConstantinopleBlock)
	}
	if isForkIncompatible(c.PetersburgBlock, newcfg.PetersburgBlock, head) {
		return newCompatError("ConstantinopleFix fork block", c.PetersburgBlock, newcfg.PetersburgBlock)
	}
	if isForkIncompatible(c.EWASMBlock, newcfg.EWASMBlock, head) {
		return newCompatError("ewasm fork block", c.EWASMBlock, newcfg.EWASMBlock)
	}
	if c.Dccs != nil && newcfg.Dccs != nil {
		if isForkIncompatible(c.Dccs.ThangLongBlock, newcfg.Dccs.ThangLongBlock, head) {
			return newCompatError("Thang Long fork block", c.Dccs.ThangLongBlock, newcfg.Dccs.ThangLongBlock)
		}
	}
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntactic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainID                                     *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158   bool
	IsByzantium, IsConstantinople, IsPetersburg bool
	IsThangLong                                 bool
}

// Rules ensures c's ChainID is not nil.
func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainID := c.ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	return Rules{
		ChainID:          new(big.Int).Set(chainID),
		IsHomestead:      c.IsHomestead(num),
		IsEIP150:         c.IsEIP150(num),
		IsEIP155:         c.IsEIP155(num),
		IsEIP158:         c.IsEIP158(num),
		IsByzantium:      c.IsByzantium(num),
		IsConstantinople: c.IsConstantinople(num),
		IsPetersburg:     c.IsPetersburg(num),
		IsThangLong:      c.IsThangLong(num),
	}
}
