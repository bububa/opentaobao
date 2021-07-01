package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* AlibabaItemOperateDownshelf
商品下架
alibaba.item.operate.downshelf

商品下架 */
func AlibabaItemOperateDownshelf(clt *core.SDKClient, req *product.AlibabaItemOperateDownshelfAPIRequest, session string) (*product.AlibabaItemOperateDownshelfAPIResponse, error) {
	var resp product.AlibabaItemOperateDownshelfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
