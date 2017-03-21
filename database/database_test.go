package database_test

import (
	. "github.com/zjohl/go-simple-database-challenge/database"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Database", func() {

	var store Store

	BeforeEach(func() {
		store = Store{
			KeyMap: make(map[string]string),
		}
	})

	Describe("#Get", func() {
		Context("when a value exists in the store", func() {

			It("retrieves that value", func() {
				store.Set("a", "1")

				Expect(store.Get("a")).To(Equal("1"))
			})
		})

		Context("when multiple values exist in the store", func() {

			It("retrieves the value associated with the given key", func() {
				store.Set("a", "1")
				store.Set("b", "2")

				Expect(store.Get("a")).To(Equal("1"))
				Expect(store.Get("b")).To(Equal("2"))
			})
		})


		Context("when a value does not exist for the given key", func() {

			It("returns an error", func() {
				_, err := store.Get("a")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("#Set", func() {
		Context("when a value exists for the given key", func() {

			It("returns the previous value", func() {
				store.Set("a", "1")

				Expect(store.Set("a", "2")).To(Equal("1"))
			})
		})
	})

	Describe("#Unset", func() {
		Context("when a value exists for the given key", func() {

			It("unsets the key", func() {
				store.Set("a", "1")
				store.Unset("a")

				val, err := store.Get("a")
				Expect(val).NotTo(Equal("1"))
				Expect(err).To(HaveOccurred())
			})

			It("returns the previous value", func() {
				store.Set("a", "1")
				val, err := store.Unset("a")

				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal("1"))
			})
		})

		Context("when a value does not exist for the given key", func() {

			It("returns an error", func() {
				_, err := store.Unset("a")
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
