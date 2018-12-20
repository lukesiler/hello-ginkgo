package api_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/lukesiler/hello-ginkgo/test/globals"
	"github.com/lukesiler/hello-ginkgo/types"
)

var _ = Describe("Root API", func() {

	BeforeEach(func() {

	})

	JustBeforeEach(func() {

	})

	Describe("HTTP server address", func() {
		It("should be initialized with address and port", func() {
			Expect(globals.HTTPServerAddress).ToNot(BeZero())

			parts := strings.Split(globals.HTTPServerAddress, ":")
			Expect(len(parts)).To(Equal(2))
		})
	})

	Describe("GET /", func() {
		It("API root w/o trailing slash should be functional and polite", func() {
			resp, err := http.Get("http://" + globals.HTTPServerAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(resp).ToNot(BeZero())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(resp.Status).To(Equal("200 OK"))

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body).To(Equal([]byte("hello!")))
		})

		It("API root w/ trailing slash should be functional and polite", func() {
			resp, err := http.Get("http://" + globals.HTTPServerAddress + "/")
			Expect(err).ToNot(HaveOccurred())
			Expect(resp).ToNot(BeZero())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(resp.Status).To(Equal("200 OK"))

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).ToNot(HaveOccurred())
			Expect(body).To(Equal([]byte("hello!")))
		})
	})

	Describe("GET /book", func() {
		It("API /book should be functional", func() {
			resp, err := http.Get("http://" + globals.HTTPServerAddress + "/book")
			Expect(err).ToNot(HaveOccurred())
			Expect(resp).ToNot(BeZero())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(resp.Status).To(Equal("200 OK"))

			defer resp.Body.Close()
			b := &types.Book{}
			err = json.NewDecoder(resp.Body).Decode(b)
			Expect(err).ToNot(HaveOccurred())
			Expect(b).ToNot(BeZero())
			Expect(b.Author).ToNot(BeZero())
			Expect(b.Author).To(Equal("Luke Siler"))
			Expect(b.Title).ToNot(BeZero())
			Expect(b.Title).To(Equal("Nothing Special"))
			Expect(b.PageCount).Should(BeNumerically(">", 0))
			Expect(b.PageCount).To(Equal(527))
		})
	})

	JustAfterEach(func() {
		if CurrentGinkgoTestDescription().Failed {
			fmt.Printf("Collecting diags just after failed test in %s\n", CurrentGinkgoTestDescription().TestText)
		}
	})
})
