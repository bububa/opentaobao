package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceitemupdate 修改商品开票信息
// alibaba.einvoice.item.update
//
// ERP通过接口将商品的开票信息同步给阿里发票平台，自动开票时将读取这些开票信息，需要联系阿里小二开通对应的权限
func Alibabaeinvoiceitemupdate(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceitemupdateAPIRequest, session string) (*einvoice.AlibabaeinvoiceitemupdateAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceitemupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
