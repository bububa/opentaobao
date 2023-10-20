package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemvipupdateschemaget vip商家编辑商品的规则获取接口
// tmall.item.vip.update.schema.get
//
// 获取vip商家编辑商品的规则
func Tmallitemvipupdateschemaget(clt *core.SDKClient, req *product.TmallitemvipupdateschemagetAPIRequest, session string) (*product.TmallitemvipupdateschemagetAPIResponse, error) {
	var resp product.TmallitemvipupdateschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
