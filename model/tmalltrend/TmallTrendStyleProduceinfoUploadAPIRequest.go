package tmalltrend

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalltrendstyleproduceinfouploadAPIRequest 款式生产信息同步API API请求
// tmall.trend.style.produceinfo.upload
//
// 款式生产信息同步至平台
type TmalltrendstyleproduceinfouploadAPIRequest struct {
	model.Params
	// 款式生产信息列表，单次同步最对1000条
	_styleProduceInfoBoList []StyleProduceInfoBo
}

// NewTmalltrendstyleproduceinfouploadRequest 初始化TmalltrendstyleproduceinfouploadAPIRequest对象
func NewTmalltrendstyleproduceinfouploadRequest() *TmalltrendstyleproduceinfouploadAPIRequest {
	return &TmalltrendstyleproduceinfouploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalltrendstyleproduceinfouploadAPIRequest) GetApiMethodName() string {
	return "tmall.trend.style.produceinfo.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalltrendstyleproduceinfouploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalltrendstyleproduceinfouploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStyleProduceInfoBoList is StyleProduceInfoBoList Setter
// 款式生产信息列表，单次同步最对1000条
func (r *TmalltrendstyleproduceinfouploadAPIRequest) SetStyleProduceInfoBoList(_styleProduceInfoBoList []StyleProduceInfoBo) error {
	r._styleProduceInfoBoList = _styleProduceInfoBoList
	r.Set("style_produce_info_bo_list", _styleProduceInfoBoList)
	return nil
}

// GetStyleProduceInfoBoList StyleProduceInfoBoList Getter
func (r TmalltrendstyleproduceinfouploadAPIRequest) GetStyleProduceInfoBoList() []StyleProduceInfoBo {
	return r._styleProduceInfoBoList
}
