package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateGet 酒店产品库rate查询
// taobao.xhotel.rate.get
//
// 酒店产品库rate查询
func TaobaoXhotelRateGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateGetAPIRequest, session string) (*xhotelitem.TaobaoXhotelRateGetAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelRateGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
