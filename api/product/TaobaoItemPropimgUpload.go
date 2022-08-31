package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoItemPropimgUpload 添加或修改属性图片
// taobao.item.propimg.upload
//
// 添加一张商品属性图片到num_iid指定的商品中 &lt;br/&gt;传入的num_iid所对应的商品必须属于当前会话的用户 &lt;br/&gt;图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的 &lt;br/&gt;商品属性图片只有享有服务的卖家（如：淘宝大卖家、订购了淘宝多图服务的卖家）才能上传 &lt;br/&gt;商品属性图片有数量和大小上的限制，最多不能超过24张（每个颜色属性都有一张）。
func TaobaoItemPropimgUpload(clt *core.SDKClient, req *product.TaobaoItemPropimgUploadAPIRequest, session string) (*product.TaobaoItemPropimgUploadAPIResponse, error) {
	var resp product.TaobaoItemPropimgUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
