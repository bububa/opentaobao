package icbuproduct

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuproduct"
)

// AlibabaIcbuProductTypeAvailableGet 商家发品类型查询
// alibaba.icbu.product.type.available.get
//
// 查询商家发品权限
func AlibabaIcbuProductTypeAvailableGet(clt *core.SDKClient, req *icbuproduct.AlibabaIcbuProductTypeAvailableGetAPIRequest, session string) (*icbuproduct.AlibabaIcbuProductTypeAvailableGetAPIResponse, error) {
	var resp icbuproduct.AlibabaIcbuProductTypeAvailableGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
