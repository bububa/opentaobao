package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelratedelete rate删除接口
// taobao.xhotel.rate.delete
//
// 酒店产品库rate删除
func Taobaoxhotelratedelete(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelratedeleteAPIRequest, session string) (*xhotelitem.TaobaoxhotelratedeleteAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelratedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
