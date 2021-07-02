package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoItemPermitCheck 发品资质校验
// taobao.item.permit.check
//
// 对淘宝商品发品、编辑前的预校验接口
func TaobaoItemPermitCheck(clt *core.SDKClient, req *product.TaobaoItemPermitCheckAPIRequest, session string) (*product.TaobaoItemPermitCheckAPIResponse, error) {
	var resp product.TaobaoItemPermitCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
