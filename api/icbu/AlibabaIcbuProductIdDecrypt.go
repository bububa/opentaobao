package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

/* AlibabaIcbuProductIdDecrypt
商品ID解密
alibaba.icbu.product.id.decrypt

对混淆的产品ID解密 */
func AlibabaIcbuProductIdDecrypt(clt *core.SDKClient, req *icbu.AlibabaIcbuProductIdDecryptAPIRequest, session string) (*icbu.AlibabaIcbuProductIdDecryptAPIResponse, error) {
	var resp icbu.AlibabaIcbuProductIdDecryptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
