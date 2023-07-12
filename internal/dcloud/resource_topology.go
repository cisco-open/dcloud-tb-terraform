// Copyright 2023 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package dcloud

import (
	"context"
	"github.com/cisco-open/dcloud-tb-go-client/tbclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTopology() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTopologyCreate,
		ReadContext:   resourceTopologyRead,
		UpdateContext: resourceTopologyUpdate,
		DeleteContext: resourceTopologyDelete,
		Schema: map[string]*schema.Schema{
			"uid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"datacenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTopologyCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	c := i.(*tbclient.Client)

	var diags diag.Diagnostics

	topology := tbclient.Topology{
		Name:        data.Get("name").(string),
		Description: data.Get("description").(string),
		Datacenter:  data.Get("datacenter").(string),
		Notes:       data.Get("notes").(string),
	}

	t, err := c.CreateTopology(topology)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(t.Uid)

	resourceTopologyRead(ctx, data, i)

	return diags

}

func resourceTopologyRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	c := i.(*tbclient.Client)

	var diags diag.Diagnostics

	t, err := c.GetTopology(data.Id())
	if err != nil {
		return handleClientError(err, data, diags)
	}

	data.Set("uid", t.Uid)
	data.Set("name", t.Name)
	data.Set("description", t.Description)
	data.Set("notes", t.Notes)
	data.Set("datacenter", t.Datacenter)
	data.Set("status", t.Status)

	return diags
}

func resourceTopologyUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	c := i.(*tbclient.Client)

	topology := tbclient.Topology{
		Uid:         data.Get("uid").(string),
		Name:        data.Get("name").(string),
		Description: data.Get("description").(string),
		Datacenter:  data.Get("datacenter").(string),
		Notes:       data.Get("notes").(string),
	}

	_, err := c.UpdateTopology(topology)
	if err != nil {
		var diags diag.Diagnostics
		return handleClientError(err, data, diags)
	}

	return resourceTopologyRead(ctx, data, i)
}

func resourceTopologyDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	c := i.(*tbclient.Client)

	if err := c.DeleteTopology(data.Id()); err != nil {
		return diag.FromErr(err)
	}

	return diag.Diagnostics{}
}
