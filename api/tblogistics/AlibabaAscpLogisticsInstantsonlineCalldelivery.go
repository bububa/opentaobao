package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Alibabaascplogisticsinstantsonlinecalldelivery 同城配送在线下单正式下单呼叫运力
// alibaba.ascp.logistics.instantsonline.calldelivery
//
// 同城配送在线下单正式下单呼叫运力
func Alibabaascplogisticsinstantsonlinecalldelivery(clt *core.SDKClient, req *tblogistics.AlibabaascplogisticsinstantsonlinecalldeliveryAPIRequest, session string) (*tblogistics.AlibabaascplogisticsinstantsonlinecalldeliveryAPIResponse, error) {
	var resp tblogistics.AlibabaascplogisticsinstantsonlinecalldeliveryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
