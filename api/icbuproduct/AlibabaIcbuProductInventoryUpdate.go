package icbuproduct

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuproduct"
)

// Alibabaicbuproductinventoryupdate icbu商品库存更新
// alibaba.icbu.product.inventory.update
//
// 更新库存信息
func Alibabaicbuproductinventoryupdate(clt *core.SDKClient, req *icbuproduct.AlibabaicbuproductinventoryupdateAPIRequest, session string) (*icbuproduct.AlibabaicbuproductinventoryupdateAPIResponse, error) {
	var resp icbuproduct.AlibabaicbuproductinventoryupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
