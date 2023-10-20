package lstfundbill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstfundbill"
)

// Alibabalsttradesellerorderfundbilllistquery 结算明细数据查询（卖家视角）
// alibaba.lst.trade.seller.order.fundbill.list.query
//
// 提供For供应商的结算接口，以交易账单维度提供开放数据，满足供应商自动化结算的诉求
func Alibabalsttradesellerorderfundbilllistquery(clt *core.SDKClient, req *lstfundbill.AlibabalsttradesellerorderfundbilllistqueryAPIRequest, session string) (*lstfundbill.AlibabalsttradesellerorderfundbilllistqueryAPIResponse, error) {
	var resp lstfundbill.AlibabalsttradesellerorderfundbilllistqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
