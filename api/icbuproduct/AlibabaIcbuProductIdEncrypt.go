package icbuproduct

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuproduct"
)

// Alibabaicbuproductidencrypt ICBU国际站商品加密接口
// alibaba.icbu.product.id.encrypt
//
// ICBU国际站，对混淆的产品ID加密。
func Alibabaicbuproductidencrypt(clt *core.SDKClient, req *icbuproduct.AlibabaicbuproductidencryptAPIRequest, session string) (*icbuproduct.AlibabaicbuproductidencryptAPIResponse, error) {
	var resp icbuproduct.AlibabaicbuproductidencryptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
