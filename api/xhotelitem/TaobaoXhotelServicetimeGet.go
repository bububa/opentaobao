package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelservicetimeget 查询实体对应的服务时间数据
// taobao.xhotel.servicetime.get
//
// 通过实体来获取对应的服务时间数据
func Taobaoxhotelservicetimeget(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelservicetimegetAPIRequest, session string) (*xhotelitem.TaobaoxhotelservicetimegetAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelservicetimegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
