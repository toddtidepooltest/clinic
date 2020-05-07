package store_test


import (
	. "github.com/onsi/ginkgo"
	"go.mongodb.org/mongo-driver/bson"

	"fmt"
	//. "github.com/onsi/ginkgo/extensions/table"
	"github.com/onsi/gomega"
	"github.com/tidepool-org/clinic/store"
	//. "github.com/onsi/gomega"
)

// NewClinic defines model for NewClinic.
type NewTestDoc struct {
	Name         *string                 `json:"name,omitempty" bson:"name,omitempty"`
	Address      *string                 `json:"address,omitempty" bson:"address,omitempty"`
	Metadata     *map[string]interface{} `json:"metadata,omitempty" bson:"metadata,omitempty"`
	PhoneNumbers *[]struct {
		Number *string `json:"number,omitempty" bson:"number,omitempty"`
		Type   *string `json:"type,omitempty" bson:"type,omitempty"`
	} `json:"phoneNumbers,omitempty" bson:"phoneNumbers,omitempty"`

}
type TestDocExtraFields struct {
	Active bool `json:"active" bson:"active"`
}
type FullTestDoc struct {
	// Embedded fields due to inline allOf schema
	Id    string `json:"_id,omitempty" bson:"_id,omitempty"`
	// Embedded struct due to allOf(#/components/schemas/ClinicianPermissions)
	NewTestDoc `bson:",inline"`
	TestDocExtraFields `bson:",inline"`
}

var (
	testName = "test"
	testAddress = "address"
)


var _ = Describe("Store Test", func() {
	Context("Database Operations", func() {
		Context("Parse", func() {
			// TODO
		})

		Context("Basic Ops", func() {

			It("should populate write data", func() {
				testDoc := FullTestDoc{NewTestDoc: NewTestDoc{Name: &testName, Address: &testAddress},
					                   TestDocExtraFields: TestDocExtraFields{Active: true}}
				_, err := mongoClient.InsertOne(store.ClinicsCollection, &testDoc)
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			})
			It("Find One After Insert", func() {
				var clinic FullTestDoc
				filter := bson.M{"name": testName}
				err := mongoClient.FindOne(store.ClinicsCollection, &filter, &clinic)
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				fmt.Printf("Clinic: %v\n", clinic)
				gomega.Expect(clinic.Address).To(gomega.Equal(&testAddress))
			})
			It("Find After Insert", func() {
				var clinic []FullTestDoc
				filter := bson.M{"name": testName}
				err := mongoClient.Find(store.ClinicsCollection, &filter, &store.DefaultPagingParams, &clinic)
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(clinic)).To(gomega.Equal(1))
				gomega.Expect(clinic[0].Address).To(gomega.Equal(&testAddress))
			})
			It("Update", func() {
				var clinic FullTestDoc
				filter := bson.M{"name": testName}

				patchObj := bson.D{
					{"$set", bson.D{{"address", "updatedAddress"}} },
				}
				err := mongoClient.UpdateOne(store.ClinicsCliniciansCollection, &filter, &patchObj)
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = mongoClient.FindOne(store.ClinicsCollection, &filter, &clinic)
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(clinic.Address).To(gomega.Equal(&testAddress))
			})
			It("Delete", func() {
				var clinic FullTestDoc
				filter := bson.M{"name": testName, "active": true}

				patchObj := bson.D{
					{"$set", bson.D{{"active", false}} },
				}
				err := mongoClient.UpdateOne(store.ClinicsCollection, &filter, &patchObj)
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				filter["active"] = false
				err = mongoClient.FindOne(store.ClinicsCollection, &filter, &clinic)
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(clinic.Address).To(gomega.Equal(&testAddress))
			})
		})

		Context("Multiple documents/Paging", func() {
		})
	})
})

