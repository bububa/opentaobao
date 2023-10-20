package xhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

// Taobaoxhoteldataserviceorderdetail 服务订单详情
// taobao.xhotel.data.service.order.detail
//
// 服务订单详情top接口构建
func Taobaoxhoteldataserviceorderdetail(clt *core.SDKClient, req *xhotel.TaobaoxhoteldataserviceorderdetailAPIRequest, session string) (*xhotel.TaobaoxhoteldataserviceorderdetailAPIResponse, error) {
	var resp xhotel.TaobaoxhoteldataserviceorderdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
