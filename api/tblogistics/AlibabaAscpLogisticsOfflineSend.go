package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Alibabaascplogisticsofflinesend 自己联系物流发货
// alibaba.ascp.logistics.offline.send
//
// 用户调用该接口可实现自己联系发货，使用该接口发货，交易订单状态会直接变成卖家已发货
func Alibabaascplogisticsofflinesend(clt *core.SDKClient, req *tblogistics.AlibabaascplogisticsofflinesendAPIRequest, session string) (*tblogistics.AlibabaascplogisticsofflinesendAPIResponse, error) {
	var resp tblogistics.AlibabaascplogisticsofflinesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
