package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbuproductiddecrypt 商品ID解密
// alibaba.icbu.product.id.decrypt
//
// 对混淆的产品ID解密
func Alibabaicbuproductiddecrypt(clt *core.SDKClient, req *icbu.AlibabaicbuproductiddecryptAPIRequest, session string) (*icbu.AlibabaicbuproductiddecryptAPIResponse, error) {
	var resp icbu.AlibabaicbuproductiddecryptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
