package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemImgUploadAPIRequest
添加商品图片 API请求
taobao.item.img.upload

添加一张商品图片到num_iid指定的商品中
传入的num_iid所对应的商品必须属于当前会话的用户
如果更新图片需要设置itemimg_id，且该itemimg_id的图片记录需要属于传入的num_iid对应的商品。如果新增图片则不用设置 。
使用taobao.item.seller.get中返回的item_imgs字段获取图片id。
商品图片有数量和大小上的限制，根据卖家享有的服务（如：卖家订购了多图服务等），商品图片数量限制不同。 */
type TaobaoItemImgUploadAPIRequest struct {
	model.Params
	// 商品图片id(如果是更新图片，则需要传该参数)
	_id int64
	// 商品数字ID，该参数必须
	_numIid int64
	// 图片序号
	_position int64
	// 商品图片内容类型:JPG;最大:3M 。支持的文件类型：jpg,jpeg,png
	_image *model.File
	// 是否将该图片设为主图,可选值:true,false;默认值:false(非主图)
	_isMajor bool
	// 是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图
	_isRectangle bool
}

// New
