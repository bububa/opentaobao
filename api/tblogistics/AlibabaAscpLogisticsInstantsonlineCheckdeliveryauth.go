package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Alibabaascplogisticsinstantsonlinecheckdeliveryauth 同城配送在线下单检查授权
// alibaba.ascp.logistics.instantsonline.checkdeliveryauth
//
// 同城配送在线下单检查授权
func Alibabaascplogisticsinstantsonlinecheckdeliveryauth(clt *core.SDKClient, req *tblogistics.AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIRequest, session string) (*tblogistics.AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIResponse, error) {
	var resp tblogistics.AlibabaascplogisticsinstantsonlinecheckdeliveryauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
