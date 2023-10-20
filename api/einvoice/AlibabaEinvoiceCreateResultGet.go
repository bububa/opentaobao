package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicecreateresultget ERP开票结果获取
// alibaba.einvoice.create.result.get
//
// ERP开票结果获取
func Alibabaeinvoicecreateresultget(clt *core.SDKClient, req *einvoice.AlibabaeinvoicecreateresultgetAPIRequest, session string) (*einvoice.AlibabaeinvoicecreateresultgetAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicecreateresultgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
