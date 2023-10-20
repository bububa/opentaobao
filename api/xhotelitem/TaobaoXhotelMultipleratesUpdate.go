package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelMultipleratesUpdate 复杂价格推送接口（批量全量）
// taobao.xhotel.multiplerates.update
//
// 批量更新复杂价格
// 涵盖了taobao.xhotel.rates.update的功能
func TaobaoXhotelMultipleratesUpdate(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelMultipleratesUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelMultipleratesUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
