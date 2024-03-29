package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceProdResultFileurlGet 发票中台-发票文件下载地址查询
// alibaba.einvoice.prod.result.fileurl.get
//
// 发票文件下载地址查询，外部ISV通过该接口可以查对应发票文件
func AlibabaEinvoiceProdResultFileurlGet(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceProdResultFileurlGetAPIRequest, resp *einvoice.AlibabaEinvoiceProdResultFileurlGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
