package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductIdDecrypt 商品ID解密
// alibaba.icbu.product.id.decrypt
//
// 对混淆的产品ID解密
func AlibabaIcbuProductIdDecrypt(clt *core.SDKClient, req *icbu.AlibabaIcbuProductIdDecryptAPIRequest, resp *icbu.AlibabaIcbuProductIdDecryptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
