package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbroomtypeDelete 民宿房源删除接口
// taobao.xhotel.bnbroomtype.delete
//
// 增加民宿房源删除接口
func TaobaoXhotelBnbroomtypeDelete(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbroomtypeDeleteAPIRequest, resp *xhotelitem.TaobaoXhotelBnbroomtypeDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
