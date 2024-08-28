package icbuproduct

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuproduct"
)

// AlibabaIcbuProductIdEncrypt ICBU国际站商品加密接口
// alibaba.icbu.product.id.encrypt
//
// ICBU国际站，对混淆的产品ID加密。
func AlibabaIcbuProductIdEncrypt(ctx context.Context, clt *core.SDKClient, req *icbuproduct.AlibabaIcbuProductIdEncryptAPIRequest, resp *icbuproduct.AlibabaIcbuProductIdEncryptAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
