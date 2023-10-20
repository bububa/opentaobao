package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Alibabaascplogisticsinstantsonlinepriorcalldelivery 同城配送在线下单预询价
// alibaba.ascp.logistics.instantsonline.priorcalldelivery
//
// 同城配送在线下单预询价
func Alibabaascplogisticsinstantsonlinepriorcalldelivery(clt *core.SDKClient, req *tblogistics.AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIRequest, session string) (*tblogistics.AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIResponse, error) {
	var resp tblogistics.AlibabaascplogisticsinstantsonlinepriorcalldeliveryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
