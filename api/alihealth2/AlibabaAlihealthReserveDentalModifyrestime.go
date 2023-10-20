package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthreservedentalmodifyrestime 修改预约时间
// alibaba.alihealth.reserve.dental.modifyrestime
//
// 修改预约时间
func Alibabaalihealthreservedentalmodifyrestime(clt *core.SDKClient, req *alihealth2.AlibabaalihealthreservedentalmodifyrestimeAPIRequest, session string) (*alihealth2.AlibabaalihealthreservedentalmodifyrestimeAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthreservedentalmodifyrestimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
