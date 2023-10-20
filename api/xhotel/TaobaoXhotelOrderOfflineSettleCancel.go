package xhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

// Taobaoxhotelorderofflinesettlecancel 线下信用住取消结账专用接口
// taobao.xhotel.order.offline.settle.cancel
//
// 线下信用住取消结账专用接口
func Taobaoxhotelorderofflinesettlecancel(clt *core.SDKClient, req *xhotel.TaobaoxhotelorderofflinesettlecancelAPIRequest, session string) (*xhotel.TaobaoxhotelorderofflinesettlecancelAPIResponse, error) {
	var resp xhotel.TaobaoxhotelorderofflinesettlecancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
