package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelServicetimeUpdate 飞猪酒店多维度服务时间维护接口
// taobao.xhotel.servicetime.update
//
// 飞猪酒店多维度服务时间维护，支持卖家维度，supplier维度，酒店维度
func TaobaoXhotelServicetimeUpdate(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelServicetimeUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelServicetimeUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
