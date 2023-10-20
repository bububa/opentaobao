package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaadgroupsbycampaignidget 批量得到推广计划下的推广单元
// taobao.simba.adgroupsbycampaignid.get
//
// 根据推广计划ID分页获取推广计划下的推广单元信息
func Taobaosimbaadgroupsbycampaignidget(clt *core.SDKClient, req *simba.TaobaosimbaadgroupsbycampaignidgetAPIRequest, session string) (*simba.TaobaosimbaadgroupsbycampaignidgetAPIResponse, error) {
	var resp simba.TaobaosimbaadgroupsbycampaignidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
