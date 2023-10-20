package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmopencustomerget 查询会员资产
// alibaba.alsc.crm.open.customer.get
//
// 查询会员身份，资产等
func Alibabaalsccrmopencustomerget(clt *core.SDKClient, req *alsc.AlibabaalsccrmopencustomergetAPIRequest, session string) (*alsc.AlibabaalsccrmopencustomergetAPIResponse, error) {
	var resp alsc.AlibabaalsccrmopencustomergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
