package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceclosereq 关闭开票失败请求（失败列表可重试）
// alibaba.einvoice.closereq
//
// 关闭失败开票请求，避免造成重复开票
func Alibabaeinvoiceclosereq(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceclosereqAPIRequest, session string) (*einvoice.AlibabaeinvoiceclosereqAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceclosereqAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
