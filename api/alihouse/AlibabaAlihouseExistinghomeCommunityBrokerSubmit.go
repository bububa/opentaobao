package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeCommunityBrokerSubmit 提交小区专家
// alibaba.alihouse.existinghome.community.broker.submit
//
// 提交小区专家
func AlibabaAlihouseExistinghomeCommunityBrokerSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
