package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthbookingreservecheckin 确认到店
// alibaba.alihealth.booking.reserve.checkin
//
// 消费医疗统一预约平台，ISV 确认到店
func Alibabaalihealthbookingreservecheckin(clt *core.SDKClient, req *alihealth2.AlibabaalihealthbookingreservecheckinAPIRequest, session string) (*alihealth2.AlibabaalihealthbookingreservecheckinAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthbookingreservecheckinAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
