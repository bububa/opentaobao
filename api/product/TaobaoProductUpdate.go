package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Taobaoproductupdate 修改一个产品，可以修改主图，不能修改子图片
// taobao.product.update
//
// 传入产品ID &lt;br/&gt;可修改字段：outer_id,binds,sale_props,name,price,desc,image &lt;br/&gt;注意：1.可以修改主图,不能修改子图片,主图最大500K,目前仅支持GIF,JPG&lt;br/&gt;      2.商城卖家产品发布24小时后不能作删除或修改操作
func Taobaoproductupdate(clt *core.SDKClient, req *product.TaobaoproductupdateAPIRequest, session string) (*product.TaobaoproductupdateAPIResponse, error) {
	var resp product.TaobaoproductupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
