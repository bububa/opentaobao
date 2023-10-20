package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelmultipleratedelete 复杂价格删除
// taobao.xhotel.multiplerate.delete
//
// 酒店产品库rate删除
func Taobaoxhotelmultipleratedelete(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelmultipleratedeleteAPIRequest, session string) (*xhotelitem.TaobaoxhotelmultipleratedeleteAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelmultipleratedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
