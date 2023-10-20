package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthbookingreserveconfirm 确认预约
// alibaba.alihealth.booking.reserve.confirm
//
// 确认预约
func Alibabaalihealthbookingreserveconfirm(clt *core.SDKClient, req *alihealth2.AlibabaalihealthbookingreserveconfirmAPIRequest, session string) (*alihealth2.AlibabaalihealthbookingreserveconfirmAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthbookingreserveconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
