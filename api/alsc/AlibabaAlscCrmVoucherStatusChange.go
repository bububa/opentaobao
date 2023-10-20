package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmvoucherstatuschange 优惠券状态更改
// alibaba.alsc.crm.voucher.status.change
//
// 核销优惠券
func Alibabaalsccrmvoucherstatuschange(clt *core.SDKClient, req *alsc.AlibabaalsccrmvoucherstatuschangeAPIRequest, session string) (*alsc.AlibabaalsccrmvoucherstatuschangeAPIResponse, error) {
	var resp alsc.AlibabaalsccrmvoucherstatuschangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
