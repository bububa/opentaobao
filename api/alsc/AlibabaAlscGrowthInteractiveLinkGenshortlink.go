package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalscgrowthinteractivelinkgenshortlink 短链接口
// alibaba.alsc.growth.interactive.link.genshortlink
//
// 短链接口
func Alibabaalscgrowthinteractivelinkgenshortlink(clt *core.SDKClient, req *alsc.AlibabaalscgrowthinteractivelinkgenshortlinkAPIRequest, session string) (*alsc.AlibabaalscgrowthinteractivelinkgenshortlinkAPIResponse, error) {
	var resp alsc.AlibabaalscgrowthinteractivelinkgenshortlinkAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
