package lstlogistics2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics2"
)

// Alibabalsttradesellerofflineorderquery 供应商-线下订单-查询接口
// alibaba.lst.trade.seller.offline.order.query
//
// 供应商线下订单数据上传后查询物流状态
func Alibabalsttradesellerofflineorderquery(clt *core.SDKClient, req *lstlogistics2.AlibabalsttradesellerofflineorderqueryAPIRequest, session string) (*lstlogistics2.AlibabalsttradesellerofflineorderqueryAPIResponse, error) {
	var resp lstlogistics2.AlibabalsttradesellerofflineorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
