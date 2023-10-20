package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbuproductupdatefield 商品按字段更新
// alibaba.icbu.product.update.field
//
// 按字段修改国际站商品，支持询盘商品和在线批发商品，支持英文商品和多语言商品
func Alibabaicbuproductupdatefield(clt *core.SDKClient, req *icbu.AlibabaicbuproductupdatefieldAPIRequest, session string) (*icbu.AlibabaicbuproductupdatefieldAPIResponse, error) {
	var resp icbu.AlibabaicbuproductupdatefieldAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
