package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeTradeEntrustSubmit 交易委托信息更新接口
// alibaba.alihouse.existinghome.trade.entrust.submit
//
// 交易委托信息更新接口
func AlibabaAlihouseExistinghomeTradeEntrustSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
