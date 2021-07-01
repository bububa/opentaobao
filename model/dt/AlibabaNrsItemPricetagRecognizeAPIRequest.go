package dt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNrsItemPricetagRecognizeAPIRequest
价签识别 API请求
alibaba.nrs.item.pricetag.recognize

商品价签识别，用于识别RT上传的竞品分析照片，返回价签内容 */
type AlibabaNrsItemPricetagRecognizeAPIRequest struct {
	model.Params
	// 图片数据
	_imgByteArr *model.File
	// 价签识别参数
	_priceTagParam *PriceTagReqParam
}

// New
