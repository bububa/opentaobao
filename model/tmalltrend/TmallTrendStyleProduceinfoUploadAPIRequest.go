package tmalltrend

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallTrendStyleProduceinfoUploadAPIRequest 款式生产信息同步API API请求
// tmall.trend.style.produceinfo.upload
//
// 款式生产信息同步至平台
type TmallTrendStyleProduceinfoUploadAPIRequest struct {
	model.Params
	// 款式生产信息列表，单次同步最对1000条
	_styleProduceInfoBoList []StyleProduceInfoBO
}

// NewTmallTrendStyleProduceinfoUploadRequest 初始化TmallTrendStyleProduceinfoUploadAPIRequest对象
func NewTmallTrendStyleProduceinfoUploadRequest() *TmallTrendStyleProduceinfoUploadAPIRequest {
	return &TmallTrendStyleProduceinfoUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTrendStyleProduceinfoUploadAPIRequest) GetApiMethodName() string {
	return "tmall.trend.style.produceinfo.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTrendStyleProduceinfoUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStyleProduceInfoBoList is StyleProduceInfoBoList Setter
// 款式生产信息列表，单次同步最对1000条
func (r *TmallTrendStyleProduceinfoUploadAPIRequest) SetStyleProduceInfoBoList(_styleProduceInfoBoList []StyleProduceInfoBO) error {
	r._styleProduceInfoBoList = _styleProduceInfoBoList
	r.Set("style_produce_info_bo_list", _styleProduceInfoBoList)
	return nil
}

// GetStyleProduceInfoBoList StyleProduceInfoBoList Getter
func (r TmallTrendStyleProduceinfoUploadAPIRequest) GetStyleProduceInfoBoList() []StyleProduceInfoBO {
	return r._styleProduceInfoBoList
}
