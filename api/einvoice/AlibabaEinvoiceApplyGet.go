package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceapplyget 开票申请数据获取接口
// alibaba.einvoice.apply.get
//
// ERP获取开票申请数据
func Alibabaeinvoiceapplyget(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceapplygetAPIRequest, session string) (*einvoice.AlibabaeinvoiceapplygetAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceapplygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
