package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthbookingreservemodify 修改预约
// alibaba.alihealth.booking.reserve.modify
//
// 消费医疗统一预约平台，取消预约
func Alibabaalihealthbookingreservemodify(clt *core.SDKClient, req *alihealth2.AlibabaalihealthbookingreservemodifyAPIRequest, session string) (*alihealth2.AlibabaalihealthbookingreservemodifyAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthbookingreservemodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
