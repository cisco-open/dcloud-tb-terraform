package dcloudtb

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
	"wwwin-github.cisco.com/pov-services/kapua-tb-go-client/tbclient"
)

func dataSourceOsFamilies() *schema.Resource {

	return &schema.Resource{
		Description: "All the OS Families available to be used in VMs",

		ReadContext: dataSourceOsFamiliesRead,

		Schema: map[string]*schema.Schema{
			"os_families": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceOsFamiliesRead(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	tb := m.(*tbclient.Client)

	osFamilies, err := tb.GetAllOsFamilies()
	if err != nil {
		return diag.FromErr(err)
	}

	osFamilyResources := make([]map[string]interface{}, len(osFamilies))
	for i, osFamily := range osFamilies {
		osFamilyResources[i] = converOsFamilyToDataResource(osFamily)
	}

	if err := data.Set("os_families", osFamilyResources); err != nil {
		return diag.FromErr(err)
	}
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diag.Diagnostics{}
}

func converOsFamilyToDataResource(osFamily tbclient.OsFamily) map[string]interface{} {
	resource := make(map[string]interface{})
	resource["id"] = osFamily.Id
	resource["name"] = osFamily.Name

	return resource
}