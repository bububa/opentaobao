package tmalltrend

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTrendStyleBasicinfoUploadAPIRequest
3D款式基本信息同步API API请求
tmall.trend.style.basicinfo.upload

3D款式基本信息同步至天猫趋势中心 */
type TmallTrendStyleBasicinfoUploadAPIRequest struct {
	model.Params
	// 款式基本信息列表，单次同步最多1000条
	_styleBasicInfoBoList []StyleBasicInfoBo
}

// NewTmallTrendStyleBasicinfoUploadRequest 初始化TmallTrendStyleBasicinfoUploadAPIRequest对象
func NewTmallTrendStyleBasicinfoUploadRequest() *TmallTrendStyleBasicinfoUploadAPIRequest {
	return &TmallTrendStyleBasicinfoUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTrendStyleBasicinfoUploadAPIRequest) GetApiMethodName() string {
	return "tmall.trend.style.basicinfo.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTrendStyleBasicinfoUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StyleBasicInfoBoList Setter
// 款式基本信息列表，单次同步最多1000条
func (r *TmallTrendStyleBasicinfoUploadAPIRequest) SetStyleBasicInfoBoList(_styleBasicInfoBoList []StyleBasicInfoBo) error {
	r._styleBasicInfoBoList = _styleBasicInfoBoList
	r.Set("style_basic_info_bo_list", _styleBasicInfoBoList)
	return nil
}

// Get StyleBasicInfoBoList Getter
func (r TmallTrendStyleBasicinfoUploadAPIRequest) GetStyleBasicInfoBoList() []StyleBasicInfoBo {
	return r._styleBasicInfoBoList
}
