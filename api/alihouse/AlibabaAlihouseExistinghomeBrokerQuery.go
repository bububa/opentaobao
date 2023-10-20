package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBrokerQuery 根据外部经纪人ID查询
// alibaba.alihouse.existinghome.broker.query
//
// 根据外部经纪人ID查询
func AlibabaAlihouseExistinghomeBrokerQuery(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBrokerQueryAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeBrokerQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
