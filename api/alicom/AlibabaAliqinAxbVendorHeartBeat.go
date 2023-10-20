package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinaxbvendorheartbeat 供应商心跳上报接口
// alibaba.aliqin.axb.vendor.heart.beat
//
// 供应商上报自己的心跳信息
func Alibabaaliqinaxbvendorheartbeat(clt *core.SDKClient, req *alicom.AlibabaaliqinaxbvendorheartbeatAPIRequest, session string) (*alicom.AlibabaaliqinaxbvendorheartbeatAPIResponse, error) {
	var resp alicom.AlibabaaliqinaxbvendorheartbeatAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
