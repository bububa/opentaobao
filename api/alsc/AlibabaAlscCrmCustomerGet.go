package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcustomerget 查询顾客详情
// alibaba.alsc.crm.customer.get
//
// 查询顾客详情
func Alibabaalsccrmcustomerget(clt *core.SDKClient, req *alsc.AlibabaalsccrmcustomergetAPIRequest, session string) (*alsc.AlibabaalsccrmcustomergetAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcustomergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
