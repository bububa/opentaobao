package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateDelete rate删除接口
// taobao.xhotel.rate.delete
//
// 酒店产品库rate删除
func TaobaoXhotelRateDelete(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateDeleteAPIRequest, resp *xhotelitem.TaobaoXhotelRateDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
