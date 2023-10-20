package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelServicetimeGet 查询实体对应的服务时间数据
// taobao.xhotel.servicetime.get
//
// 通过实体来获取对应的服务时间数据
func TaobaoXhotelServicetimeGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelServicetimeGetAPIRequest, resp *xhotelitem.TaobaoXhotelServicetimeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
