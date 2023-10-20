package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthImriskQuery 问诊质控接口
// alibaba.alihealth.imrisk.query
//
// 阿里健康的问诊质控接口
func AlibabaAlihealthImriskQuery(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthImriskQueryAPIRequest, resp *alihealth2.AlibabaAlihealthImriskQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
