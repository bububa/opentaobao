package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// Alibabaicbuquotationpost 供应商提交报价
// alibaba.icbu.quotation.post
//
// 供应商对RFQ进行报价
func Alibabaicbuquotationpost(clt *core.SDKClient, req *icburfq.AlibabaicbuquotationpostAPIRequest, session string) (*icburfq.AlibabaicbuquotationpostAPIResponse, error) {
	var resp icburfq.AlibabaicbuquotationpostAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
