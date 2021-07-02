package dt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNrsItemPricetagRecognizeAPIRequest 价签识别 API请求
// alibaba.nrs.item.pricetag.recognize
//
// 商品价签识别，用于识别RT上传的竞品分析照片，返回价签内容
type AlibabaNrsItemPricetagRecognizeAPIRequest struct {
	model.Params
	// 图片数据
	_imgByteArr *model.File
	// 价签识别参数
	_priceTagParam *PriceTagReqParam
}

// NewAlibabaNrsItemPricetagRecognizeRequest 初始化AlibabaNrsItemPricetagRecognizeAPIRequest对象
func NewAlibabaNrsItemPricetagRecognizeRequest() *AlibabaNrsItemPricetagRecognizeAPIRequest {
	return &AlibabaNrsItemPricetagRecognizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNrsItemPricetagRecognizeAPIRequest) GetApiMethodName() string {
	return "alibaba.nrs.item.pricetag.recognize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNrsItemPricetagRecognizeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetImgByteArr is ImgByteArr Setter
// 图片数据
func (r *AlibabaNrsItemPricetagRecognizeAPIRequest) SetImgByteArr(_imgByteArr *model.File) error {
	r._imgByteArr = _imgByteArr
	r.Set("img_byte_arr", _imgByteArr)
	return nil
}

// GetImgByteArr ImgByteArr Getter
func (r AlibabaNrsItemPricetagRecognizeAPIRequest) GetImgByteArr() *model.File {
	return r._imgByteArr
}

// SetPriceTagParam is PriceTagParam Setter
// 价签识别参数
func (r *AlibabaNrsItemPricetagRecognizeAPIRequest) SetPriceTagParam(_priceTagParam *PriceTagReqParam) error {
	r._priceTagParam = _priceTagParam
	r.Set("price_tag_param", _priceTagParam)
	return nil
}

// GetPriceTagParam PriceTagParam Getter
func (r AlibabaNrsItemPricetagRecognizeAPIRequest) GetPriceTagParam() *PriceTagReqParam {
	return r._priceTagParam
}
