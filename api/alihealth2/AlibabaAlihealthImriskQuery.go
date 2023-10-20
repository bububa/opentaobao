package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthImriskQuery 问诊质控接口
// alibaba.alihealth.imrisk.query
//
// 阿里健康的问诊质控接口
func AlibabaAlihealthImriskQuery(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthImriskQueryAPIRequest, session string) (*alihealth2.AlibabaAlihealthImriskQueryAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthImriskQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
