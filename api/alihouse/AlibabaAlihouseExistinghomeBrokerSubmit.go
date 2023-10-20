package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBrokerSubmit 提交经纪人信息
// alibaba.alihouse.existinghome.broker.submit
//
// 提交经纪人信息
func AlibabaAlihouseExistinghomeBrokerSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeBrokerSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
