package xhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

// Taobaoxhotelorderhotelsignquery 获取直连酒店（客栈）签约上线进度信息
// taobao.xhotel.order.hotelsign.query
//
// 获取直连酒店（客栈）签约上线进度信息
func Taobaoxhotelorderhotelsignquery(clt *core.SDKClient, req *xhotel.TaobaoxhotelorderhotelsignqueryAPIRequest, session string) (*xhotel.TaobaoxhotelorderhotelsignqueryAPIResponse, error) {
	var resp xhotel.TaobaoxhotelorderhotelsignqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
