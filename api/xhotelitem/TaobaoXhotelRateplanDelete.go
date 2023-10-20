package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateplanDelete 价格计划rateplan删除
// taobao.xhotel.rateplan.delete
//
// 酒店产品库rateplan删除，同时删除级联的rate，请谨慎使用
func TaobaoXhotelRateplanDelete(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateplanDeleteAPIRequest, resp *xhotelitem.TaobaoXhotelRateplanDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
