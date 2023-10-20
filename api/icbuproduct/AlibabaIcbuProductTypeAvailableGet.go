package icbuproduct

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuproduct"
)

// Alibabaicbuproducttypeavailableget 商家发品类型查询
// alibaba.icbu.product.type.available.get
//
// 查询商家发品权限
func Alibabaicbuproducttypeavailableget(clt *core.SDKClient, req *icbuproduct.AlibabaicbuproducttypeavailablegetAPIRequest, session string) (*icbuproduct.AlibabaicbuproducttypeavailablegetAPIResponse, error) {
	var resp icbuproduct.AlibabaicbuproducttypeavailablegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
