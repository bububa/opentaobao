package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeTradeEntrustSubmit 交易委托信息更新接口
// alibaba.alihouse.existinghome.trade.entrust.submit
//
// 交易委托信息更新接口
func AlibabaAlihouseExistinghomeTradeEntrustSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
