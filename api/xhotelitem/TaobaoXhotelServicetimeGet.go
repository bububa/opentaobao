package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelServicetimeGet 查询实体对应的服务时间数据
// taobao.xhotel.servicetime.get
//
// 通过实体来获取对应的服务时间数据
func TaobaoXhotelServicetimeGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelServicetimeGetAPIRequest, session string) (*xhotelitem.TaobaoXhotelServicetimeGetAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelServicetimeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
