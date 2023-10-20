package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelget 酒店查询接口
// taobao.xhotel.get
//
// 酒店查询接口
func Taobaoxhotelget(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelgetAPIRequest, session string) (*xhotelitem.TaobaoxhotelgetAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
