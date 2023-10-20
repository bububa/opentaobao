package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderPrePayAction 服务商预付款完成接口
// alibaba.idle.tender.pre.pay.action
//
// 服务商预付款完成接口
func AlibabaIdleTenderPrePayAction(clt *core.SDKClient, req *idle.AlibabaIdleTenderPrePayActionAPIRequest, session string) (*idle.AlibabaIdleTenderPrePayActionAPIResponse, error) {
	var resp idle.AlibabaIdleTenderPrePayActionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
