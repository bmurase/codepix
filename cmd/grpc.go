package cmd

import (
	"os"

	"github.com/bmurase/codepix/application/grpc"
	"github.com/bmurase/codepix/infrastructure/db"
	"github.com/spf13/cobra"
)

var portNumber int

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		grpc.StartGrpcServer(database, 50051)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&portNumber, "port", "p", 50051, "gRPC server port")
}
