package aquasec

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAquasecRolesDatasource(t *testing.T) {

	if isSaasEnv() {
		t.Skip("Skipping prem user test because its on saas env")
	}
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAquasecRolesDataSource(),
				Check:  testAccCheckAquasecRolesDataSourceExists("data.aquasec_roles.testroles"),
			},
		},
	})
}

func testAccCheckAquasecRolesDataSource() string {
	return `
	data "aquasec_roles" "testroles" {}
	`

}

func testAccCheckAquasecRolesDataSourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return NewNotFoundErrorf("%s in state", n)
		}

		if rs.Primary.ID == "" {
			return NewNotFoundErrorf("Id for %s in state", n)
		}

		return nil
	}
}
