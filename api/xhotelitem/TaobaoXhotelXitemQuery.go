package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelXitemQuery 查询 x 元素
// taobao.xhotel.xitem.query
//
// 查询 x 元素
func TaobaoXhotelXitemQuery(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelXitemQueryAPIRequest, resp *xhotelitem.TaobaoXhotelXitemQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
