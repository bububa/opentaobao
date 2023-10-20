package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeTradeEntrustSubmit 交易委托信息更新接口
// alibaba.alihouse.existinghome.trade.entrust.submit
//
// 交易委托信息更新接口
func AlibabaAlihouseExistinghomeTradeEntrustSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
