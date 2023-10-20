package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Alibabaascplogisticsinstantsonlinedeliveryorderget 同城配送在线下单获取配送单
// alibaba.ascp.logistics.instantsonline.deliveryorder.get
//
// 同城配送在线下单获取配送单
func Alibabaascplogisticsinstantsonlinedeliveryorderget(clt *core.SDKClient, req *tblogistics.AlibabaascplogisticsinstantsonlinedeliveryordergetAPIRequest, session string) (*tblogistics.AlibabaascplogisticsinstantsonlinedeliveryordergetAPIResponse, error) {
	var resp tblogistics.AlibabaascplogisticsinstantsonlinedeliveryordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
