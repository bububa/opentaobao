package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthbookingreservecancel 取消预约
// alibaba.alihealth.booking.reserve.cancel
//
// 消费医疗统一预约平台，ISV取消预约
func Alibabaalihealthbookingreservecancel(clt *core.SDKClient, req *alihealth2.AlibabaalihealthbookingreservecancelAPIRequest, session string) (*alihealth2.AlibabaalihealthbookingreservecancelAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthbookingreservecancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
