package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalscorderorderupload 订单回流
// alibaba.alsc.order.order.upload
//
// 第三方订单回流
func Alibabaalscorderorderupload(clt *core.SDKClient, req *alsc.AlibabaalscorderorderuploadAPIRequest, session string) (*alsc.AlibabaalscorderorderuploadAPIResponse, error) {
	var resp alsc.AlibabaalscorderorderuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
