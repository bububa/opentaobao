package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* AlibabaItemOperateUpshelf
商品上架
alibaba.item.operate.upshelf

商品上架 */
func AlibabaItemOperateUpshelf(clt *core.SDKClient, req *product.AlibabaItemOperateUpshelfAPIRequest, session string) (*product.AlibabaItemOperateUpshelfAPIResponse, error) {
	var resp product.AlibabaItemOperateUpshelfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
