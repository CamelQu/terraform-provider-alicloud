package alicloud

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alikafka"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/terraform-providers/terraform-provider-alicloud/alicloud/connectivity"
)

func init() {
	resource.AddTestSweepers("alicloud_alikafka_consumer_group", &resource.Sweeper{
		Name: "alicloud_alikafka_consumer_group",
		F:    testSweepAlikafkaConsumerGroup,
	})
}

func testSweepAlikafkaConsumerGroup(region string) error {
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		return WrapErrorf(err, "error getting Alicloud client.")
	}
	client := rawClient.(*connectivity.AliyunClient)
	alikafkaService := AlikafkaService{client}

	prefixes := []string{
		"tf-testAcc",
		"tf_testacc",
	}

	instanceListReq := alikafka.CreateGetInstanceListRequest()
	instanceListReq.RegionId = defaultRegionToTest

	raw, err := alikafkaService.client.WithAlikafkaClient(func(alikafkaClient *alikafka.Client) (interface{}, error) {
		return alikafkaClient.GetInstanceList(instanceListReq)
	})
	if err != nil {
		log.Printf("[ERROR] Failed to retrieve alikafka instance in service list: %s", err)
	}

	instanceListResp, _ := raw.(*alikafka.GetInstanceListResponse)

	var instanceIds []string
	for _, v := range instanceListResp.InstanceList.InstanceVO {
		instanceIds = append(instanceIds, v.InstanceId)
	}

	for _, instanceId := range instanceIds {

		// Control the consumer group list request rate.
		time.Sleep(time.Duration(400) * time.Millisecond)

		request := alikafka.CreateGetConsumerListRequest()
		request.InstanceId = instanceId
		request.RegionId = defaultRegionToTest

		raw, err := alikafkaService.client.WithAlikafkaClient(func(alikafkaClient *alikafka.Client) (interface{}, error) {
			return alikafkaClient.GetConsumerList(request)
		})
		if err != nil {
			log.Printf("[ERROR] Failed to retrieve alikafka consumer groups on instance (%s): %s", instanceId, err)
			continue
		}

		consumerListResp, _ := raw.(*alikafka.GetConsumerListResponse)
		consumers := consumerListResp.ConsumerList.ConsumerVO
		for _, v := range consumers {
			name := v.ConsumerId
			skip := true
			for _, prefix := range prefixes {
				if strings.HasPrefix(strings.ToLower(name), strings.ToLower(prefix)) {
					skip = false
					break
				}
			}
			if skip {
				log.Printf("[INFO] Skipping alikafka consumer id: %s ", name)
				continue
			}
			log.Printf("[INFO] delete alikafka consumer group: %s ", name)

			// Control the consumer group delete rate
			time.Sleep(time.Duration(400) * time.Millisecond)

			request := alikafka.CreateDeleteConsumerGroupRequest()
			request.InstanceId = instanceId
			request.ConsumerId = v.ConsumerId
			request.RegionId = defaultRegionToTest

			_, err := alikafkaService.client.WithAlikafkaClient(func(alikafkaClient *alikafka.Client) (interface{}, error) {
				return alikafkaClient.DeleteConsumerGroup(request)
			})
			if err != nil {
				log.Printf("[ERROR] Failed to delete alikafka consumer group (%s): %s", name, err)
			}
		}
	}

	return nil
}

func TestAccAlicloudAlikafkaConsumerGroup_basic(t *testing.T) {

	var v *alikafka.ConsumerVO
	resourceId := "alicloud_alikafka_consumer_group.default"
	ra := resourceAttrInit(resourceId, alikafkaConsumerGroupBasicMap)
	serviceFunc := func() interface{} {
		return &AlikafkaService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandInt()
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testacc-alikafkaconsumerbasic%v", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceAlikafkaConsumerGroupConfigDependence)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, connectivity.AlikafkaSupportedRegions)
			testAccPreCheck(t)
			testAccPreCheckWithAlikafkaInstanceSetting(t)
		},
		// module name
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_id": os.Getenv("ALICLOUD_INSTANCE_ID"),
					"consumer_id": "${var.consumer_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"consumer_id": fmt.Sprintf("tf-testacc-alikafkaconsumerbasic%v", rand),
					}),
				),
			},

			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})

}

func TestAccAlicloudAlikafkaConsumerGroup_multi(t *testing.T) {

	var v *alikafka.ConsumerVO
	resourceId := "alicloud_alikafka_consumer_group.default.4"
	ra := resourceAttrInit(resourceId, alikafkaConsumerGroupBasicMap)
	serviceFunc := func() interface{} {
		return &AlikafkaService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandInt()
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testacc-alikafkaconsumerbasic%v", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceAlikafkaConsumerGroupConfigDependence)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, connectivity.AlikafkaSupportedRegions)
			testAccPreCheck(t)
			testAccPreCheckWithAlikafkaInstanceSetting(t)
		},
		// module name
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"count":       "5",
					"instance_id": os.Getenv("ALICLOUD_INSTANCE_ID"),
					"consumer_id": "${var.consumer_id}-${count.index}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"consumer_id": fmt.Sprintf("tf-testacc-alikafkaconsumerbasic%v-4", rand),
					}),
				),
			},
		},
	})

}

func resourceAlikafkaConsumerGroupConfigDependence(name string) string {
	return fmt.Sprintf(`
		variable "consumer_id" {
 			default = "%v"
		}
		`, name)
}

var alikafkaConsumerGroupBasicMap = map[string]string{
	"consumer_id": "${var.consumer_id}",
}
