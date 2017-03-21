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
			OccurrencesMap: make(map[string]int),
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

	Describe("#NumEqualTo", func() {
		Context("when there are no stored instances of the provided value", func() {

			It("returns 0", func() {
				Expect(store.NumEqualTo("1")).To(Equal(0))
			})
		})

		Context("when there is a stored instance of the provided value", func() {

			It("returns the number of occurances of the value", func() {
				store.Set("a", "1")

				Expect(store.NumEqualTo("1")).To(Equal(1))
			})
		})

		Context("when there are many stored instances of the provided value", func() {

			It("returns the number of occurances of the value", func() {
				store.Set("a", "1")
				store.Set("b", "1")
				store.Set("d", "1")

				Expect(store.NumEqualTo("1")).To(Equal(3))
			})
		})
	})
})
