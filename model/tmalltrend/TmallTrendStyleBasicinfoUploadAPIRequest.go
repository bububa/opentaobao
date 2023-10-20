package tmalltrend

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalltrendstylebasicinfouploadAPIRequest 3D款式基本信息同步API API请求
// tmall.trend.style.basicinfo.upload
//
// 3D款式基本信息同步至天猫趋势中心
type TmalltrendstylebasicinfouploadAPIRequest struct {
	model.Params
	// 款式基本信息列表，单次同步最多1000条
	_styleBasicInfoBoList []StyleBasicInfoBo
}

// NewTmalltrendstylebasicinfouploadRequest 初始化TmalltrendstylebasicinfouploadAPIRequest对象
func NewTmalltrendstylebasicinfouploadRequest() *TmalltrendstylebasicinfouploadAPIRequest {
	return &TmalltrendstylebasicinfouploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalltrendstylebasicinfouploadAPIRequest) GetApiMethodName() string {
	return "tmall.trend.style.basicinfo.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalltrendstylebasicinfouploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalltrendstylebasicinfouploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStyleBasicInfoBoList is StyleBasicInfoBoList Setter
// 款式基本信息列表，单次同步最多1000条
func (r *TmalltrendstylebasicinfouploadAPIRequest) SetStyleBasicInfoBoList(_styleBasicInfoBoList []StyleBasicInfoBo) error {
	r._styleBasicInfoBoList = _styleBasicInfoBoList
	r.Set("style_basic_info_bo_list", _styleBasicInfoBoList)
	return nil
}

// GetStyleBasicInfoBoList StyleBasicInfoBoList Getter
func (r TmalltrendstylebasicinfouploadAPIRequest) GetStyleBasicInfoBoList() []StyleBasicInfoBo {
	return r._styleBasicInfoBoList
}
