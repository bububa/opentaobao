package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicedeductget 发票扣减的接口
// alibaba.einvoice.deduct.get
//
// 获取历史发票扣减量、每日发票扣减量的接口
func Alibabaeinvoicedeductget(clt *core.SDKClient, req *einvoice.AlibabaeinvoicedeductgetAPIRequest, session string) (*einvoice.AlibabaeinvoicedeductgetAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicedeductgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
