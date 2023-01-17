package tmalltrend

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallTrendStyleBindinfoUploadAPIRequest 趋势词&款式绑定信息同步API API请求
// tmall.trend.style.bindinfo.upload
//
// 趋势词&amp;款式(服饰行业)绑定信息同步至平台
type TmallTrendStyleBindinfoUploadAPIRequest struct {
	model.Params
	// 趋势词&款式绑定信息列表，一次最多1000条
	_trendStyleBindInfoBoList []TrendStyleBindInfoBO
}

// NewTmallTrendStyleBindinfoUploadRequest 初始化TmallTrendStyleBindinfoUploadAPIRequest对象
func NewTmallTrendStyleBindinfoUploadRequest() *TmallTrendStyleBindinfoUploadAPIRequest {
	return &TmallTrendStyleBindinfoUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTrendStyleBindinfoUploadAPIRequest) GetApiMethodName() string {
	return "tmall.trend.style.bindinfo.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTrendStyleBindinfoUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTrendStyleBindinfoUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrendStyleBindInfoBoList is TrendStyleBindInfoBoList Setter
// 趋势词&amp;款式绑定信息列表，一次最多1000条
func (r *TmallTrendStyleBindinfoUploadAPIRequest) SetTrendStyleBindInfoBoList(_trendStyleBindInfoBoList []TrendStyleBindInfoBO) error {
	r._trendStyleBindInfoBoList = _trendStyleBindInfoBoList
	r.Set("trend_style_bind_info_bo_list", _trendStyleBindInfoBoList)
	return nil
}

// GetTrendStyleBindInfoBoList TrendStyleBindInfoBoList Getter
func (r TmallTrendStyleBindinfoUploadAPIRequest) GetTrendStyleBindInfoBoList() []TrendStyleBindInfoBO {
	return r._trendStyleBindInfoBoList
}
