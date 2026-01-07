package entity_test

import (
	"se-final-lab-exam/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestGeneralCase(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	books := entity.Books{
		Title: "Harry Potter",
		Price: 249.0,
		Code: "BK001474",
	}
	
	ok, err := govalidator.ValidateStruct(books)
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())
}