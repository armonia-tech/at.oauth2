package generates_test

import (
	"context"
	"testing"
	"time"

	oauth2 "github.com/armonia-tech/at.oauth2/v4"
	generates "github.com/armonia-tech/at.oauth2/v4/generates"
	models "github.com/armonia-tech/at.oauth2/v4/models"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAuthorize(t *testing.T) {
	Convey("Test Authorize Generate", t, func() {
		data := &oauth2.GenerateBasic{
			Client: &models.Client{
				ID:     "123456",
				Secret: "123456",
			},
			UserID:   "000000",
			CreateAt: time.Now(),
		}
		gen := generates.NewAuthorizeGenerate()
		code, err := gen.Token(context.Background(), data)
		So(err, ShouldBeNil)
		So(code, ShouldNotBeEmpty)
		Println("\nAuthorize Code:" + code)
	})
}
