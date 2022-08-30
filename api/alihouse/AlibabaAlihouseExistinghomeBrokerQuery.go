package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBrokerQuery 根据外部经纪人ID查询
// alibaba.alihouse.existinghome.broker.query
//
// 根据外部经纪人ID查询
func AlibabaAlihouseExistinghomeBrokerQuery(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBrokerQueryAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeBrokerQueryAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeBrokerQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
