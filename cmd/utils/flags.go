package utils

import (
	//"crypto/ecdsa"
	//"fmt"
	//"io"
	//"io/ioutil"
	//"math/big"
	//"os"
	//"path/filepath"
	//"strconv"
	//"strings"
	//"text/tabwriter"
	//"text/template"
	//"time"

	//"github.com/RWAValueRouter/FastMulThreshold-DSA/accounts"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/accounts/keystore"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/common"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/common/fdlimit"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/consensus"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/consensus/clique"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/consensus/ethash"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/core"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/core/rawdb"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/core/vm"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/crypto"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/eth"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/eth/downloader"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/eth/gasprice"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/ethdb"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/ethstats"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/graphql"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/internal/ethapi"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/internal/flags"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/les"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/log"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/metrics"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/metrics/exp"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/metrics/influxdb"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/miner"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/node"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/p2p"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/p2p/discv5"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/p2p/enode"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/p2p/nat"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/p2p/netutil"
	//"github.com/RWAValueRouter/FastMulThreshold-DSA/params"
	//pcsclite "github.com/gballet/go-libpcsclite"
	cli "gopkg.in/urfave/cli.v1"
)

// MigrateFlags sets the global flag from a local flag when it's set.
// This is a temporary function used for migrating old command/flags to the
// new format.
//
// e.g. geth account new --keystore /tmp/mykeystore --lightkdf
//
// is equivalent after calling this method with:
//
// geth --keystore /tmp/mykeystore --lightkdf account new
//
// This allows the use of the existing configuration functionality.
// When all flags are migrated this function can be removed and the existing
// configuration functionality must be changed that is uses local flags
func MigrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		for _, name := range ctx.FlagNames() {
			if ctx.IsSet(name) {
			    err := ctx.GlobalSet(name, ctx.String(name))
			    if err != nil {
				return err
			    }
			}
		}
		return action(ctx)
	}
}
