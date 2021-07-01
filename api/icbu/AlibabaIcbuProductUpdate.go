package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

/* AlibabaIcbuProductUpdate
修改商品
alibaba.icbu.product.update

修改国际站商品，支持询盘商品和在线批发商品，支持英文商品和多语言商品 */
func AlibabaIcbuProductUpdate(clt *core.SDKClient, req *icbu.AlibabaIcbuProductUpdateAPIRequest, session string) (*icbu.AlibabaIcbuProductUpdateAPIResponse, error) {
	var resp icbu.AlibabaIcbuProductUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
