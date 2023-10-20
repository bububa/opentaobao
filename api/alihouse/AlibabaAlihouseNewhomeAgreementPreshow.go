package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeagreementpreshow 预览地址获取接口
// alibaba.alihouse.newhome.agreement.preshow
//
// 预览地址获取接口
func Alibabaalihousenewhomeagreementpreshow(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeagreementpreshowAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeagreementpreshowAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeagreementpreshowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
