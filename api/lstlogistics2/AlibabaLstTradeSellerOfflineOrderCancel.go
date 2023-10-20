package lstlogistics2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics2"
)

// Alibabalsttradesellerofflineordercancel 供应商-线下订单-取消接口
// alibaba.lst.trade.seller.offline.order.cancel
//
// 供应商线下订单数据上传之后取消
func Alibabalsttradesellerofflineordercancel(clt *core.SDKClient, req *lstlogistics2.AlibabalsttradesellerofflineordercancelAPIRequest, session string) (*lstlogistics2.AlibabalsttradesellerofflineordercancelAPIResponse, error) {
	var resp lstlogistics2.AlibabalsttradesellerofflineordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
