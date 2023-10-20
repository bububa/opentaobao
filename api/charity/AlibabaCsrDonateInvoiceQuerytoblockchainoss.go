package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacsrdonateinvoicequerytoblockchainoss 触发odps任务离线查询公益宝贝开票对账明细
// alibaba.csr.donate.invoice.querytoblockchainoss
//
// 提供给蚂蚁链上公益团队，用于触发odps任务离线查询公益宝贝开票对账明细
func Alibabacsrdonateinvoicequerytoblockchainoss(clt *core.SDKClient, req *charity.AlibabacsrdonateinvoicequerytoblockchainossAPIRequest, session string) (*charity.AlibabacsrdonateinvoicequerytoblockchainossAPIResponse, error) {
	var resp charity.AlibabacsrdonateinvoicequerytoblockchainossAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
