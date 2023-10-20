package xhoteloffline

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhoteloffline"
)

// Taobaoxhotelorderalipayfacecreate 信用住支付创建接口
// taobao.xhotel.order.alipayface.create
//
// 用于创建一笔信用住支付，主要应用场景是线下信用住
func Taobaoxhotelorderalipayfacecreate(clt *core.SDKClient, req *xhoteloffline.TaobaoxhotelorderalipayfacecreateAPIRequest, session string) (*xhoteloffline.TaobaoxhotelorderalipayfacecreateAPIResponse, error) {
	var resp xhoteloffline.TaobaoxhotelorderalipayfacecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
