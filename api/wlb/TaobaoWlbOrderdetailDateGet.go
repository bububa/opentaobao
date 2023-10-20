package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Taobaowlborderdetaildateget 按照日期范围查询物流订单详情
// taobao.wlb.orderdetail.date.get
//
// 外部ERP可通过该接口查询一段时间内的物流宝订单，以及订单详情
func Taobaowlborderdetaildateget(clt *core.SDKClient, req *wlb.TaobaowlborderdetaildategetAPIRequest, session string) (*wlb.TaobaowlborderdetaildategetAPIResponse, error) {
	var resp wlb.TaobaowlborderdetaildategetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
