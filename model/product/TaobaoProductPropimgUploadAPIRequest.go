package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoProductPropimgUploadAPIRequest
上传单张产品属性图片，如果需要传多张，可调多次 API请求
taobao.product.propimg.upload

传入产品ID <br/>传入props,目前仅支持颜色属性.调用taobao.itemprops.get.v2取得颜色属性pid,<br/>再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串; <br/>传入图片内容 <br/>注意：图片最大为2M,只支持JPG,GIF,如果需要传多张，可调多次 */
type TaobaoProductPropimgUploadAPIRequest struct {
	model.Params
	// 产品属性图片ID
	_id int64
	// 产品ID.Product的id
	_productId int64
	// 属性串.目前仅支持颜色属性.调用taobao.itemprops.get获取类目属性,取得颜色属性pid,再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串;
	_props string
	// 图片内容.图片最大为2M,只支持JPG,GIF.
	_image *model.File
	// 图片序号
	_position int64
}

// New
