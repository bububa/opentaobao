package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelGetEntityByTag 根据标签查询实体
// taobao.xhotel.get.entity.by.tag
//
// 根据标签查询实体
func TaobaoXhotelGetEntityByTag(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelGetEntityByTagAPIRequest, resp *xhotelitem.TaobaoXhotelGetEntityByTagAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
