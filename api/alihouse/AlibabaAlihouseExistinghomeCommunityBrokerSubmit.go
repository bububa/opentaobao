package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeCommunityBrokerSubmit 提交小区专家
// alibaba.alihouse.existinghome.community.broker.submit
//
// 提交小区专家
func AlibabaAlihouseExistinghomeCommunityBrokerSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeCommunityBrokerSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
