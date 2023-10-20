package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalscchudatemplatesend 本地生活触达模板消息发送接口
// alibaba.alsc.chuda.template.send
//
// 允许三方小程序通过该api发送本地生活触达消息
func Alibabaalscchudatemplatesend(clt *core.SDKClient, req *alsc.AlibabaalscchudatemplatesendAPIRequest, session string) (*alsc.AlibabaalscchudatemplatesendAPIResponse, error) {
	var resp alsc.AlibabaalscchudatemplatesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
