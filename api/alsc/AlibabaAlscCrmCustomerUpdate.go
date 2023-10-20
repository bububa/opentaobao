package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcustomerupdate 更新顾客信息
// alibaba.alsc.crm.customer.update
//
// 更新顾客信息
func Alibabaalsccrmcustomerupdate(clt *core.SDKClient, req *alsc.AlibabaalsccrmcustomerupdateAPIRequest, session string) (*alsc.AlibabaalsccrmcustomerupdateAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcustomerupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
