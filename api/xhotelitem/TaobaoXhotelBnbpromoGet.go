package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbpromoGet 民宿查询营销活动
// taobao.xhotel.bnbpromo.get
//
// 民宿查询营销活动
func TaobaoXhotelBnbpromoGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbpromoGetAPIRequest, session string) (*xhotelitem.TaobaoXhotelBnbpromoGetAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelBnbpromoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
