package xhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

/* TaobaoXhotelDataServiceOrderDetail
服务订单详情
taobao.xhotel.data.service.order.detail

服务订单详情top接口构建 */
func TaobaoXhotelDataServiceOrderDetail(clt *core.SDKClient, req *xhotel.TaobaoXhotelDataServiceOrderDetailAPIRequest, session string) (*xhotel.TaobaoXhotelDataServiceOrderDetailAPIResponse, error) {
	var resp xhotel.TaobaoXhotelDataServiceOrderDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
