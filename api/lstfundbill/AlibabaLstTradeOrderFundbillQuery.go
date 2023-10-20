package lstfundbill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstfundbill"
)

// AlibabaLstTradeOrderFundbillQuery 结算明细数据查询（品牌商视角）
// alibaba.lst.trade.order.fundbill.query
//
// 按照指定日期提供交易账单维度的结算明细数据，比供应商工作台上的结算账单还多一些数据项。
func AlibabaLstTradeOrderFundbillQuery(clt *core.SDKClient, req *lstfundbill.AlibabaLstTradeOrderFundbillQueryAPIRequest, session string) (*lstfundbill.AlibabaLstTradeOrderFundbillQueryAPIResponse, error) {
	var resp lstfundbill.AlibabaLstTradeOrderFundbillQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
