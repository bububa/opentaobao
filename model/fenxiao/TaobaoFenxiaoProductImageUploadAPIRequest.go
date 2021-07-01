package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductImageUploadAPIRequest
产品图片上传 API请求
taobao.fenxiao.product.image.upload

产品主图图片空间相对路径或绝对路径添加或更新，或者是图片上传。如果指定位置的图片已存在，则覆盖原有信息。如果位置为1,自动设为主图；如果位置为0，表示属性图片 */
type TaobaoFenxiaoProductImageUploadAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// 产品主图图片空间相对路径或绝对路径
	_picPath string
	// 产品图片
	_image *model.File
	// 图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图
	_position int64
	// properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
	_properties string
}

// New
