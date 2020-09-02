// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package unifi

import (
	"context"
	"fmt"
)

// just to fix compile issues with the import
var (
	_ fmt.Formatter
	_ context.Context
)

type ChannelPlan struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	ApBlacklistedChannels   []ChannelPlan_ApBlacklistedChannels   `json:"ap_blacklisted_channels,omitempty"`
	ConfSource              string                                `json:"conf_source,omitempty"` // manual|radio-ai
	Coupling                []ChannelPlan_Coupling                `json:"coupling,omitempty"`
	Date                    string                                `json:"date"` // ^$|^(20[0-9]{2}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9])Z?$
	Fitness                 float64                               `json:"fitness,omitempty"`
	Note                    string                                `json:"note,omitempty"`  // .{0,1024}
	Radio                   string                                `json:"radio,omitempty"` // na|ng|ng\+na
	RadioTable              []ChannelPlan_RadioTable              `json:"radio_table,omitempty"`
	Satisfaction            float64                               `json:"satisfaction,omitempty"`
	SatisfactionTable       []ChannelPlan_SatisfactionTable       `json:"satisfaction_table,omitempty"`
	SiteBlacklistedChannels []ChannelPlan_SiteBlacklistedChannels `json:"site_blacklisted_channels,omitempty"`
}

type ChannelPlan_ApBlacklistedChannels struct {
	Channel   int    `json:"channel,omitempty"`   // 36|38|40|42|44|46|48|52|56|60|64|100|104|108|112|116|120|124|128|132|136|140|144|149|153|157|161|165|183|184|185|187|188|189|192|196
	MAC       string `json:"mac,omitempty"`       // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	Timestamp int    `json:"timestamp,omitempty"` // [1-9][0-9]{12}
}

type ChannelPlan_Coupling struct {
	Rssi   int    `json:"rssi,omitempty"`
	Source string `json:"source,omitempty"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2}).*$
	Target string `json:"target,omitempty"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2}).*$
}

type ChannelPlan_RadioTable struct {
	BackupChannel string `json:"backup_channel,omitempty"` // [0-9]|[1][0-4]|16|34|36|38|40|42|44|46|48|52|56|60|64|100|104|108|112|116|120|124|128|132|136|140|144|149|153|157|161|165|183|184|185|187|188|189|192|196|auto
	Channel       string `json:"channel,omitempty"`        // [0-9]|[1][0-4]|16|34|36|38|40|42|44|46|48|52|56|60|64|100|104|108|112|116|120|124|128|132|136|140|144|149|153|157|161|165|183|184|185|187|188|189|192|196|auto
	DeviceMAC     string `json:"device_mac,omitempty"`     // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	Name          string `json:"name,omitempty"`           // [a-z]*[0-9]*
	TxPower       string `json:"tx_power,omitempty"`       // [\d]+|auto
	TxPowerMode   string `json:"tx_power_mode,omitempty"`  // auto|medium|high|low|custom
	Width         int    `json:"width,omitempty"`          // 20|40|80|160
}

type ChannelPlan_SatisfactionTable struct {
	DeviceMAC    string  `json:"device_mac,omitempty"` // ^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$
	Satisfaction float64 `json:"satisfaction,omitempty"`
}

type ChannelPlan_SiteBlacklistedChannels struct {
	Channel   int `json:"channel,omitempty"`   // 36|38|40|42|44|46|48|52|56|60|64|100|104|108|112|116|120|124|128|132|136|140|144|149|153|157|161|165|183|184|185|187|188|189|192|196
	Timestamp int `json:"timestamp,omitempty"` // [1-9][0-9]{12}
}

func (c *Client) listChannelPlan(ctx context.Context, site string) ([]ChannelPlan, error) {
	var respBody struct {
		Meta meta          `json:"meta"`
		Data []ChannelPlan `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/rest/channelplan", site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

func (c *Client) getChannelPlan(ctx context.Context, site, id string) (*ChannelPlan, error) {
	var respBody struct {
		Meta meta          `json:"meta"`
		Data []ChannelPlan `json:"data"`
	}

	err := c.do(ctx, "GET", fmt.Sprintf("s/%s/rest/channelplan/%s", site, id), nil, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	d := respBody.Data[0]
	return &d, nil
}

func (c *Client) deleteChannelPlan(ctx context.Context, site, id string) error {
	err := c.do(ctx, "DELETE", fmt.Sprintf("s/%s/rest/channelplan/%s", site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) createChannelPlan(ctx context.Context, site string, d *ChannelPlan) (*ChannelPlan, error) {
	var respBody struct {
		Meta meta          `json:"meta"`
		Data []ChannelPlan `json:"data"`
	}

	err := c.do(ctx, "POST", fmt.Sprintf("s/%s/rest/channelplan", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}

func (c *Client) updateChannelPlan(ctx context.Context, site string, d *ChannelPlan) (*ChannelPlan, error) {
	var respBody struct {
		Meta meta          `json:"meta"`
		Data []ChannelPlan `json:"data"`
	}

	err := c.do(ctx, "PUT", fmt.Sprintf("s/%s/rest/channelplan/%s", site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	if len(respBody.Data) != 1 {
		return nil, &NotFoundError{}
	}

	new := respBody.Data[0]

	return &new, nil
}