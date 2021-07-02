package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateplanAdd 酒店产品库rateplan添加
// taobao.xhotel.rateplan.add
//
// 酒店产品库rateplan
func TaobaoXhotelRateplanAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateplanAddAPIRequest, session string) (*xhotelitem.TaobaoXhotelRateplanAddAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelRateplanAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
