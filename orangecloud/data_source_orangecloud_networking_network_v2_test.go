package orangecloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

// PASS
func TestAccOrangeCloudNetworkingNetworkV2DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccOrangeCloudNetworkingNetworkV2DataSource_network,
			},
			resource.TestStep{
				Config: testAccOrangeCloudNetworkingNetworkV2DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkingNetworkV2DataSourceID("data.orangecloud_networking_network_v2.net"),
					resource.TestCheckResourceAttr(
						"data.orangecloud_networking_network_v2.net", "name", "tf_test_network"),
					resource.TestCheckResourceAttr(
						"data.orangecloud_networking_network_v2.net", "admin_state_up", "true"),
				),
			},
		},
	})
}

// PASS
func TestAccOrangeCloudNetworkingNetworkV2DataSource_subnet(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccOrangeCloudNetworkingNetworkV2DataSource_network,
			},
			resource.TestStep{
				Config: testAccOrangeCloudNetworkingNetworkV2DataSource_subnet,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkingNetworkV2DataSourceID("data.orangecloud_networking_network_v2.net"),
					resource.TestCheckResourceAttr(
						"data.orangecloud_networking_network_v2.net", "name", "tf_test_network"),
					resource.TestCheckResourceAttr(
						"data.orangecloud_networking_network_v2.net", "admin_state_up", "true"),
				),
			},
		},
	})
}

// PASS
func TestAccOrangeCloudNetworkingNetworkV2DataSource_networkID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccOrangeCloudNetworkingNetworkV2DataSource_network,
			},
			resource.TestStep{
				Config: testAccOrangeCloudNetworkingNetworkV2DataSource_networkID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkingNetworkV2DataSourceID("data.orangecloud_networking_network_v2.net"),
					resource.TestCheckResourceAttr(
						"data.orangecloud_networking_network_v2.net", "name", "tf_test_network"),
					resource.TestCheckResourceAttr(
						"data.orangecloud_networking_network_v2.net", "admin_state_up", "true"),
				),
			},
		},
	})
}

func testAccCheckNetworkingNetworkV2DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find network data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Network data source ID not set")
		}

		return nil
	}
}

const testAccOrangeCloudNetworkingNetworkV2DataSource_network = `
resource "orangecloud_networking_network_v2" "net" {
        name = "tf_test_network"
        admin_state_up = "true"
}

resource "orangecloud_networking_subnet_v2" "subnet" {
  name = "tf_test_subnet"
  cidr = "192.168.199.0/24"
  no_gateway = true
  network_id = "${orangecloud_networking_network_v2.net.id}"
}
`

var testAccOrangeCloudNetworkingNetworkV2DataSource_basic = fmt.Sprintf(`
%s

data "orangecloud_networking_network_v2" "net" {
	name = "${orangecloud_networking_network_v2.net.name}"
}
`, testAccOrangeCloudNetworkingNetworkV2DataSource_network)

var testAccOrangeCloudNetworkingNetworkV2DataSource_subnet = fmt.Sprintf(`
%s

data "orangecloud_networking_network_v2" "net" {
	matching_subnet_cidr = "${orangecloud_networking_subnet_v2.subnet.cidr}"
}
`, testAccOrangeCloudNetworkingNetworkV2DataSource_network)

var testAccOrangeCloudNetworkingNetworkV2DataSource_networkID = fmt.Sprintf(`
%s

data "orangecloud_networking_network_v2" "net" {
	network_id = "${orangecloud_networking_network_v2.net.id}"
}
`, testAccOrangeCloudNetworkingNetworkV2DataSource_network)
