package icbudropshipping

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// Alibabaorderlogisticstrackingget 阿里巴巴订单物流轨迹查询
// alibaba.order.logistics.tracking.get
//
// 阿里巴巴订单物流轨迹查询
func Alibabaorderlogisticstrackingget(clt *core.SDKClient, req *icbudropshipping.AlibabaorderlogisticstrackinggetAPIRequest, session string) (*icbudropshipping.AlibabaorderlogisticstrackinggetAPIResponse, error) {
	var resp icbudropshipping.AlibabaorderlogisticstrackinggetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
