package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthtracecodesellercoderelationcodeantiactive 解除码的关联关系
// alibaba.alihealth.tracecodeseller.code.relation.codeantiactive
//
// 解除码的关联关系
func Alibabaalihealthtracecodesellercoderelationcodeantiactive(clt *core.SDKClient, req *alihealth2.AlibabaalihealthtracecodesellercoderelationcodeantiactiveAPIRequest, session string) (*alihealth2.AlibabaalihealthtracecodesellercoderelationcodeantiactiveAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthtracecodesellercoderelationcodeantiactiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
