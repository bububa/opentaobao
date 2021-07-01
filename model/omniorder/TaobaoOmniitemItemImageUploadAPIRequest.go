package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemItemImageUploadAPIRequest
全渠道商品上传图片 API请求
taobao.omniitem.item.image.upload

全渠道商品上传图片 */
type TaobaoOmniitemItemImageUploadAPIRequest struct {
	model.Params
	// 商品图片信息，允许png、jpg、gif图片格式,3M以内
	_img *model.File
	// 条形码
	_barCode string
	// 商品ID，若填入商品ID则以商品ID为准，否则以outerId/barCode为准
	_itemId int64
	// 商品outerId
	_outerId string
	// 是否为主图
	_major bool
	// 图片顺序
	_position int64
}

// New
