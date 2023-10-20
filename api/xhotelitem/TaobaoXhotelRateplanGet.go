package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateplanGet 价格计划rateplan查询
// taobao.xhotel.rateplan.get
//
// 酒店产品库rateplan查询
func TaobaoXhotelRateplanGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateplanGetAPIRequest, session string) (*xhotelitem.TaobaoXhotelRateplanGetAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelRateplanGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
