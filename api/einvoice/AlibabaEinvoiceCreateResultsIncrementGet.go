package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicecreateresultsincrementget ERP增量开票结果获取
// alibaba.einvoice.create.results.increment.get
//
// 增量开票结果获取
func Alibabaeinvoicecreateresultsincrementget(clt *core.SDKClient, req *einvoice.AlibabaeinvoicecreateresultsincrementgetAPIRequest, session string) (*einvoice.AlibabaeinvoicecreateresultsincrementgetAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicecreateresultsincrementgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
