package api_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"

	"github.com/lukesiler/hello-ginkgo/api"
	"github.com/lukesiler/hello-ginkgo/test/globals"
)

var (
	addr       string
	logger     *log.Logger
	httpServer *http.Server
	serverNode int
)

func TestApi(t *testing.T) {
	// sole connection point between ginkgo and gomega
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Suite")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	fmt.Printf("SyncedBeforeSuite running on node %v of %v\n", config.GinkgoConfig.ParallelNode, config.GinkgoConfig.ParallelTotal)
	serverNode = config.GinkgoConfig.ParallelNode

	addr = api.ServeAddress()
	// passing URL state to test specs this way...likely a better way
	globals.HTTPServerAddress = addr

	logger = log.New(os.Stdout, "", 0)
	httpServer, gracefulStopChan, err := api.ServeAPI()
	if err != nil {
		logger.Panic(err)
	}

	go func() {
		sig := <-gracefulStopChan

		logger.Println(fmt.Sprintf("Captured %v. Shutting down HTTP server.", sig))
		httpServer.Shutdown(context.Background())
		logger.Println("Server gracefully shutdown")
	}()

	return []byte(addr)
}, func(data []byte) {
	// ensure addr is initialized and known on all ginkgo nodes not just initiator
	addr = string(data)
})

// var _ = BeforeSuite(func() {
// 	var err error

// 	Expect(err).NotTo(HaveOccurred())

// 	Expect(err).NotTo(HaveOccurred())
// })

var _ = AfterSuite(func() {
	if serverNode == 0 || config.GinkgoConfig.ParallelNode != serverNode || httpServer == nil {
		return
	}

	logger.Println("Suite complete. Shutting down HTTP server.")
	httpServer.Shutdown(context.Background())
	logger.Println("Server gracefully shutdown")
})
