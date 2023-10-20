package xhoteloffline

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhoteloffline"
)

// Taobaoxhotelorderofflinesettleput 线下信用住结账专用接口
// taobao.xhotel.order.offline.settle.put
//
// 线下信用住结账专用接口
func Taobaoxhotelorderofflinesettleput(clt *core.SDKClient, req *xhoteloffline.TaobaoxhotelorderofflinesettleputAPIRequest, session string) (*xhoteloffline.TaobaoxhotelorderofflinesettleputAPIResponse, error) {
	var resp xhoteloffline.TaobaoxhotelorderofflinesettleputAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
