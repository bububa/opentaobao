package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelentityconfig 飞猪商品各实体通用配置
// taobao.xhotel.entity.config
//
// 飞猪商品各实体通用配置服务
func Taobaoxhotelentityconfig(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelentityconfigAPIRequest, session string) (*xhotelitem.TaobaoxhotelentityconfigAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelentityconfigAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
