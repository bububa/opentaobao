package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// Alibabaascplogisticsconsignmodify 修改物流公司和运单号
// alibaba.ascp.logistics.consign.modify
//
// 修改物流公司和运单号
func Alibabaascplogisticsconsignmodify(clt *core.SDKClient, req *tblogistics.AlibabaascplogisticsconsignmodifyAPIRequest, session string) (*tblogistics.AlibabaascplogisticsconsignmodifyAPIResponse, error) {
	var resp tblogistics.AlibabaascplogisticsconsignmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
