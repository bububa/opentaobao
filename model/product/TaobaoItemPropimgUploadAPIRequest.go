package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemPropimgUploadAPIRequest
添加或修改属性图片 API请求
taobao.item.propimg.upload

添加一张商品属性图片到num_iid指定的商品中 <br/>传入的num_iid所对应的商品必须属于当前会话的用户 <br/>图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的 <br/>商品属性图片只有享有服务的卖家（如：淘宝大卖家、订购了淘宝多图服务的卖家）才能上传 <br/>商品属性图片有数量和大小上的限制，最多不能超过24张（每个颜色属性都有一张）。 */
type TaobaoItemPropimgUploadAPIRequest struct {
	model.Params
	// 商品数字ID，必选
	_numIid int64
	// 属性列表。调用taobao.itemprops.get获取类目属性，属性必须是颜色属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid。
	_properties string
	// 属性图片内容。类型:JPG,GIF;图片大小不超过:3M
	_image *model.File
	// 属性图片ID。如果是新增不需要填写
	_id int64
	// 图片位置
	_position int64
}

// New
