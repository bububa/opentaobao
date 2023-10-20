package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcustomervoucherlist 获取顾客优惠券列表
// alibaba.alsc.crm.customer.voucher.list
//
// 获取顾客优惠券列表
func Alibabaalsccrmcustomervoucherlist(clt *core.SDKClient, req *alsc.AlibabaalsccrmcustomervoucherlistAPIRequest, session string) (*alsc.AlibabaalsccrmcustomervoucherlistAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcustomervoucherlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
