package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallProductSpecPicUploadAPIRequest
上传产品规格认证图片 API请求
tmall.product.spec.pic.upload

上传指定类型的产品规格认证文件，并返回存有上传成功图片url的产品规格对象 */
type TmallProductSpecPicUploadAPIRequest struct {
	model.Params
	// 上传的认证图片的认证类型<br>1：代表产品包装正面图<br>2：代表完整产品资质<br>3：代表产品包装反面图<br>4：代表产品包装侧面图<br>5：代表产品包装条形码特写<br>6：代表特殊用途化妆品批准文号<br>7：代表3C认证图标<br>
	_certifyType int64
	// 上传的认证图片文件
	_certifyPic *model.File
}

// New
