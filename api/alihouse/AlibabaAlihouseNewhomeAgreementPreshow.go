package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeAgreementPreshow 预览地址获取接口
// alibaba.alihouse.newhome.agreement.preshow
//
// 预览地址获取接口
func AlibabaAlihouseNewhomeAgreementPreshow(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeAgreementPreshowAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeAgreementPreshowAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
