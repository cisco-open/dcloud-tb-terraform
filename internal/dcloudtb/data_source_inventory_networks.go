package dcloudtb

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
	"wwwin-github.cisco.com/pov-services/kapua-tb-go-client/tbclient"
)

func dataSourceInventoryNetworks() *schema.Resource {

	return &schema.Resource{
		Description: "All the inventory networks available to be used in a topology",

		ReadContext: dataSourceInventoryNetworksRead,

		Schema: map[string]*schema.Schema{
			"topology_uid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"inventory_networks": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceInventoryNetworksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	tb := m.(*tbclient.Client)

	topologyUid := d.Get("topology_uid").(string)

	inventoryNetworks, err := tb.GetAllInventoryNetworks(topologyUid)
	if err != nil {
		return diag.FromErr(err)
	}

	inventoryNetworkResources := make([]map[string]interface{}, len(inventoryNetworks))

	for i, inventoryNetwork := range inventoryNetworks {
		inventoryNetworkResources[i] = convertInventoryNetworkToDataResource(inventoryNetwork, topologyUid)
	}

	if err := d.Set("inventory_networks", inventoryNetworkResources); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diag.Diagnostics{}
}

func convertInventoryNetworkToDataResource(inventoryNetwork tbclient.InventoryNetwork, topologyUid string) map[string]interface{} {
	resource := make(map[string]interface{})
	resource["id"] = inventoryNetwork.Id
	resource["type"] = inventoryNetwork.Type
	resource["subnet"] = inventoryNetwork.Subnet

	return resource
}
