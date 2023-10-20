package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Alibabaascplogisticsinstantsonlinecanceldelivery 同城配送在线下单取消下单取消呼叫的运力
// alibaba.ascp.logistics.instantsonline.canceldelivery
//
// 同城配送在线下单取消下单取消呼叫的运力
func Alibabaascplogisticsinstantsonlinecanceldelivery(clt *core.SDKClient, req *tblogistics.AlibabaascplogisticsinstantsonlinecanceldeliveryAPIRequest, session string) (*tblogistics.AlibabaascplogisticsinstantsonlinecanceldeliveryAPIResponse, error) {
	var resp tblogistics.AlibabaascplogisticsinstantsonlinecanceldeliveryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
