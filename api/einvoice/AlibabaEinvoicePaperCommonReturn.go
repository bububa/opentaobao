package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicepapercommonreturn 纸票通用回传接口
// alibaba.einvoice.paper.common.return
//
// 纸票通用回传接口（打印回传、注册回传等），只返回成功or失败
func Alibabaeinvoicepapercommonreturn(clt *core.SDKClient, req *einvoice.AlibabaeinvoicepapercommonreturnAPIRequest, session string) (*einvoice.AlibabaeinvoicepapercommonreturnAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicepapercommonreturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
