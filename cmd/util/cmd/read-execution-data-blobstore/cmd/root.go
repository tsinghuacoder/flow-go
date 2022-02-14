package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	badger "github.com/ipfs/go-ds-badger2"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	"github.com/onflow/flow-go/module/blobs"
	"github.com/onflow/flow-go/network"
	"github.com/onflow/flow-go/network/mocknetwork"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/mock"
)

var (
	flagBlobstoreDir string
)

var rootCmd = &cobra.Command{
	Use:   "execution-data-blobstore",
	Short: "interact with execution data blobstore",
}

var RootCmd = rootCmd

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
}

func initBlobservice() network.BlobService {
	ds, err := badger.NewDatastore(flagBlobstoreDir, &badger.DefaultOptions)

	if err != nil {
		log.Fatal().Err(err).Msg("could not init badger datastore")
	}

	defer ds.Close()

	bstore := blockstore.NewBlockstore(ds)

	blockService := blockservice.New(bstore, nil)

	defer blockService.Close()

	bs := new(mocknetwork.BlobService)

	bs.
		On("GetBlobs", mock.Anything, mock.AnythingOfType("[]cid.Cid")).
		Return(func(ctx context.Context, ks []cid.Cid) <-chan blobs.Blob {
			return blockService.GetBlocks(ctx, ks)
		})

	return bs
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&flagBlobstoreDir, "blobstore-dir", "d", "./execution_data_blobstore", "directory to the execution data blobstore")
	_ = rootCmd.MarkPersistentFlagRequired("blobstore-dir")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AutomaticEnv()
}
