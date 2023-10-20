package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelGet 酒店查询接口
// taobao.xhotel.get
//
// 酒店查询接口
func TaobaoXhotelGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelGetAPIRequest, resp *xhotelitem.TaobaoXhotelGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
