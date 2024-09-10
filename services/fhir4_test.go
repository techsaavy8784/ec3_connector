package services

import (
	//log "github.com/sirupsen/logrus"
	//. "github.com/smartystreets/goconvey/convey"

	"fmt"
	//"os"
	"testing"

	//"time"

	//github.com/davecgh/go-spew/spew"
	//log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

//const pid = "Tbt3KuCY0B5PSrJvCu2j-PlK.aiHsu2xUjUM8bWpetXoB"

//const ordercode = "8310-5"
//const baseurl = "https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/"

const pid = "4342009"
const baseurl = "https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d/"

func TestQuery(t *testing.T) {
	fmt.Printf("Test run a FHIR query")
	c := New(baseurl)
	Convey("Run a query", t, func() {
		data, err := c.Query("Patient/12724066")
		So(err, ShouldBeNil)
		So(data, ShouldNotBeNil)
		//fmt.Printf("Patient returned: %s\n", spew.Sdump(data))
	})
}

func TestDocumentReferenceQuery(t *testing.T) {
	fmt.Printf("\n\n\n\nFhir4Test:38  --  Test run a FHIR query")
	c := New(baseurl)
	Convey("Run a query", t, func() {
		data, err := c.Query("DocumentReference?patient=12724066")
		So(err, ShouldBeNil)
		So(data, ShouldNotBeNil)
		//fmt.Printf("Patient returned: %s\n", spew.Sdump(data))
	})
}
func TestGetFhirPdf(t *testing.T) {
	Convey("Subject: GetFhirPdf", t, func() {
		fmt.Printf("TestGetFhirPDF\n")
		url := fmt.Sprintf("%s%s%s", baseurl, "/Binary/XR-", "197198634")
		//url := "https://fhir-open.cerner.com/dstu2/ec2458f2-1e24-41c8-b71b-0e701af7583d/Patient?-pageContext=2d61b0b7-805d-4fd5-bb1d-a111f942f7a5&-pageDirection=NEXT"
		c := New(baseurl)

		data, err := c.GetFhir(url)
		So(err, ShouldBeNil)
		So(data, ShouldNotBeNil)
	})
}

func TestGetFHIR(t *testing.T) {
	Convey("Subject: GetFHIR", t, func() {
		fmt.Printf("TestGetFHIR\n")
		fmt.Printf("baseURL: %s\n", baseurl)
		url := "https://fhir-open.cerner.com/dstu2/ec2458f2-1e24-41c8-b71b-0e701af7583d/Patient?family=smart&given=na"
		c := New(baseurl)
		//url = "Patient?family=smart&given=na"
		data, err := c.GetFhir(url)
		So(err, ShouldBeNil)
		So(data, ShouldNotBeNil)
	})
}

// func TestGetDocumentImage(t *testing.T) {

// 	Convey("Get the imag of a document", t, func() {
// 		//m.DeleteDocuments(session.CacheName)
// 		Convey("Authorized with a valid session", func() {
// 			url := "https://fhir-open.cerner.com/dstu2/ec2458f2-1e24-41c8-b71b-0e701af7583d/Binary/XR-197293272"
// 			c := New(baseurl)
// 			//bytes, err := c.GetFhir(url, "application/pdf")
// 			bytes, err := c.GetDiagnosticPDF(url)
// 			So(err, ShouldBeNil)
// 			So(bytes, ShouldNotBeNil)
// 			fmt.Printf("Writting the Datafile\n")
// 			if err := os.WriteFile("./debbie.data", bytes, 0666); err != nil {
// 				log.Fatal(err)
// 			}

// 		})
// 	})
// }

// func TestPatientDocumentReference(t *testing.T) {
// 	c := New(baseurl)
// 	Convey("Get GetpatientDocumentReference", t, func() {
// 		//https://fhir-open.sandboxcerner.com/dstu2/0b8a0111-e8e6-4c26-a91c-5069cbc6b1ca/DiagnosticReport?patient=1316020&_count=10
// 		data, err := c.GetPatientDiagnosticReports("12724066")
// 		So(err, ShouldBeNil)
// 		So(data, ShouldNotBeNil)
// 		fmt.Printf("Document: %s\n", spew.Sdump(data))
// 	})

// 	//data, err := c.GetDocumentReference(pid)

// }

// func TestPatientDiagnosticReports(t *testing.T) {
// 	c := New(baseurl)
// 	Convey("Get GetpatientDiagnosticReoprts", t, func() {
// 		//https://fhir-open.sandboxcerner.com/dstu2/0b8a0111-e8e6-4c26-a91c-5069cbc6b1ca/DiagnosticReport?patient=1316020&_count=10
// 		data, err := c.GetPatientDiagnosticReports("12724066")
// 		So(err, ShouldBeNil)
// 		So(data, ShouldNotBeNil)
// 		fmt.Printf("Document: %s\n", spew.Sdump(data))
// 	})

// 	//data, err := c.GetDocumentReference(pid)

// }
// func TestPatientDiagnosticReport(t *testing.T) {
// 	c := New(baseurl)
// 	// Provide the PatientID , DocumentID and DateRange of the document
// 	// R2 cerner/Epic Does not support direct query by DocumentID. Cerner does not have an R4 DiagRept

// 	Convey("Get GetpatientDiagnosticReoprts", t, func() {
// 		//https://fhir-open.sandboxcerner.com/dstu2/0b8a0111-e8e6-4c26-a91c-5069cbc6b1ca/DiagnosticReport?patient=1316020&_count=10
// 		data, err := c.GetPatientDiagnosticReports("4342009")
// 		So(err, ShouldBeNil)
// 		So(data, ShouldNotBeNil)
// 		//data, err := c.GetDocumentReference(pid)

// 		//fmt.Printf("Document: %s\n", spew.Sdump(data))
// 		// if data.Total == 0 {
// 		// 	t.Error("Expected > 0 got 0")
// 		// }
// 	})
// }

// func TestDocument(t *testing.T) {
// 	c := New(baseurl)
// 	data, err := c.GetDocumentReference("12724066")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("Document: %s\n", spew.Sdump((data)))
// 	// if data.Total == 0 {
// 	// 	t.Error("Expected > 0 got 0")
// 	// }
// }

// // func TestCondition(t *testing.T) {
// // 	c := New(baseurl)
// // 	data, err := c.GetCondition(pid)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if data.Total == 0 {
// // 		t.Error("Expected > 0 got 0")
// // 	}
// // }

// // func TestProcedure(t *testing.T) {
// // 	c := New(baseurl)
// // 	data, err := c.GetProcedure(pid)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if data.Total == 0 {
// // 		t.Error("Expected > 0 got 0")
// // 	}
// // }

// // func TestMedication(t *testing.T) {
// // 	c := New(baseurl)
// // 	data, err := c.GetMedication(pid)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if data.Total == 0 {
// // 		t.Error("Expected > 0 got 0")
// // 	}
// // }

// // func TestObservation(t *testing.T) {
// // 	c := New(baseurl)
// // 	data, err := c.GetObservation(pid, ordercode)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if data.Total == 0 {
// // 		t.Error("Expected > 0 got 0")
// // 	}
// // }

// // func TestImmunization(t *testing.T) {
// // 	c := New(baseurl)
// // 	data, err := c.GetImmunization(pid)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if data.Total == 0 {
// // 		t.Error("Expected > 0 got 0")
// // 	}
// // }

// // func TestAllergy(t *testing.T) {
// // 	c := New(baseurl)
// // 	data, err := c.GetAllergyIntolerence(pid)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if data.Total == 0 {
// // 		t.Error("Expected > 0 got 0")
// // 	}
// // }

// // func TestFamilyHx(t *testing.T) {
// // 	c := New(baseurl)
// // 	data, err := c.GetFamilyMemberHistory(pid)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if data.Total == 0 {
// // 		t.Error("Expected > 0 got 0")
// // 	}
// // }
