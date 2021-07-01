package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoProductImgUploadAPIRequest
上传单张产品非主图，如果需要传多张，可调多次 API请求
taobao.product.img.upload

1.传入产品ID <br/>2.传入图片内容 <br/>注意：图片最大为500K,只支持JPG,GIF格式,如果需要传多张，可调多次 */
type TaobaoProductImgUploadAPIRequest struct {
	model.Params
	// 产品图片ID.修改图片时需要传入
	_id int64
	// 产品ID.Product的id
	_productId int64
	// 图片内容.图片最大为500K,只支持JPG,GIF格式.
	_image *model.File
	// 图片序号
	_position int64
	// 是否将该图片设为主图.可选值:true,false;默认值:false.
	_isMajor bool
}

// New
