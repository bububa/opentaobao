package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

/* AlibabaIcbuQuotationPost
供应商提交报价
alibaba.icbu.quotation.post

供应商对RFQ进行报价 */
func AlibabaIcbuQuotationPost(clt *core.SDKClient, req *icburfq.AlibabaIcbuQuotationPostAPIRequest, session string) (*icburfq.AlibabaIcbuQuotationPostAPIResponse, error) {
	var resp icburfq.AlibabaIcbuQuotationPostAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
