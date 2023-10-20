package icbuproduct

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuproduct"
)

// AlibabaIcbuProductIdEncrypt ICBU国际站商品加密接口
// alibaba.icbu.product.id.encrypt
//
// ICBU国际站，对混淆的产品ID加密。
func AlibabaIcbuProductIdEncrypt(clt *core.SDKClient, req *icbuproduct.AlibabaIcbuProductIdEncryptAPIRequest, resp *icbuproduct.AlibabaIcbuProductIdEncryptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
