package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthBookingReserveTreat 确认就诊
// alibaba.alihealth.booking.reserve.treat
//
// 统一预约平台，ISV确认就诊
func AlibabaAlihealthBookingReserveTreat(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBookingReserveTreatAPIRequest, session string) (*alihealth2.AlibabaAlihealthBookingReserveTreatAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthBookingReserveTreatAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
