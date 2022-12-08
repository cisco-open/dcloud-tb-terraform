package dcloudtb

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
	"wwwin-github.cisco.com/pov-services/kapua-tb-go-client/tbclient"
)

func dataSourceNetworks() *schema.Resource {

	return &schema.Resource{
		Description: "All the networks currently in a given topology",

		ReadContext: dataSourceNetworksRead,

		Schema: map[string]*schema.Schema{
			"topology_uid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"inventory_network_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"inventory_network_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"inventory_network_subnet": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"topology_uid": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	tb := m.(*tbclient.Client)

	topologyUid := d.Get("topology_uid").(string)

	networks, err := tb.GetAllNetworks(topologyUid)
	if err != nil {
		return diag.FromErr(err)
	}

	networkResources := make([]map[string]interface{}, len(networks))

	for i, network := range networks {
		networkResources[i] = convertNetworkToDataResource(network)
	}

	if err := d.Set("networks", networkResources); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diag.Diagnostics{}
}

func convertNetworkToDataResource(network tbclient.Network) map[string]interface{} {
	resource := make(map[string]interface{})
	resource["uid"] = network.Uid
	resource["name"] = network.Name
	resource["description"] = network.Description
	resource["inventory_network_id"] = network.InventoryNetwork.Id
	resource["inventory_network_type"] = network.InventoryNetwork.Type
	resource["inventory_network_subnet"] = network.InventoryNetwork.Subnet
	resource["topology_uid"] = network.Topology.Uid

	return resource
}
