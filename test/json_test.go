package api_test

import (
	"encoding/json"
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/lukesiler/hello-ginkgo/types"
)

var _ = Describe("Root API", func() {
	var (
		book       types.Book
		err        error
		jsonString string
	)

	BeforeEach(func() {
		jsonString = `{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "page_count":1488
		}`
	})

	JustBeforeEach(func() {
		// runs after all BeforeEach instances but before It - BeforeEach are from outside to in and then same for JustBeforeEach
		err = json.Unmarshal([]byte(jsonString), &book)
	})

	Describe("loading from JSON", func() {
		Context("when the JSON parses succesfully", func() {
			if false {
				Skip("a special way to skip tests and exit closure")
			}

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should populate the fields correctly", func() {
				Expect(book.Title).To(Equal("Les Miserables"))
				Expect(book.Author).To(Equal("Victor Hugo"))
				Expect(book.PageCount).To(Equal(1488))
			})
		})

		Context("when the JSON fails to parse", func() {
			BeforeEach(func() {
				book = types.Book{}
				jsonString = `{
                    "title":"Les Miserables",
                    "author":"Victor Hugo",
                    "page_count":1488oops
                }`
			})

			It("unmarshal should error", func() {
				Expect(err).To(HaveOccurred())
			})

			It("should return the zero-value for the book", func() {
				Expect(book).To(BeZero())
			})
		})
	})

	Describe("Extracting the author's last name", func() {
		It("should correctly identify and return the last name", func() {
			Expect(strings.Split(book.Author, " ")[1]).To(Equal("Hugo"))
		})
	})

	JustAfterEach(func() {
		if CurrentGinkgoTestDescription().Failed {
			fmt.Printf("Collecting diags just after failed test in %s\n", CurrentGinkgoTestDescription().TestText)
			fmt.Printf("Actual book was %v\n", book)
		}
	})
})
