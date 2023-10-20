package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcardbindcustomer 卡号绑定顾客
// alibaba.alsc.crm.card.bindcustomer
//
// 为卡号绑定顾客
func Alibabaalsccrmcardbindcustomer(clt *core.SDKClient, req *alsc.AlibabaalsccrmcardbindcustomerAPIRequest, session string) (*alsc.AlibabaalsccrmcardbindcustomerAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcardbindcustomerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
