package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmvouchertemplatelist 获取优惠券模版列表
// alibaba.alsc.crm.voucher.template.list
//
// 获取优惠券模版列表
func Alibabaalsccrmvouchertemplatelist(clt *core.SDKClient, req *alsc.AlibabaalsccrmvouchertemplatelistAPIRequest, session string) (*alsc.AlibabaalsccrmvouchertemplatelistAPIResponse, error) {
	var resp alsc.AlibabaalsccrmvouchertemplatelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
