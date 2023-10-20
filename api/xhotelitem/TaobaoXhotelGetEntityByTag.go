package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelgetentitybytag 根据标签查询实体
// taobao.xhotel.get.entity.by.tag
//
// 根据标签查询实体
func Taobaoxhotelgetentitybytag(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelgetentitybytagAPIRequest, session string) (*xhotelitem.TaobaoxhotelgetentitybytagAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelgetentitybytagAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
