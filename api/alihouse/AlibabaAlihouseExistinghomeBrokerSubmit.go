package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBrokerSubmit 提交经纪人信息
// alibaba.alihouse.existinghome.broker.submit
//
// 提交经纪人信息
func AlibabaAlihouseExistinghomeBrokerSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeBrokerSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
