package alitripdivisions

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripplatformdivisionsgetdivisionbynameAPIRequest 根据中文名称与行政区划级别查询行政区划数据 API请求
// alitrip.platform.divisions.getdivisionbyname
//
// 根据中文名称与行政区划级别查询行政区划数据
type AlitripplatformdivisionsgetdivisionbynameAPIRequest struct {
	model.Params
	// 行政区划名称
	_name string
	// 行政区划级别ALL(0, "全部"),  	CONTINENT(1, "大洲"),  	COUNTRY(2, "国家"),  	PROVINCE(3, "省份"),  	CITY(4, "城市"),  	DISTRICT(5, "区县"),  	STREET(6, "街道")
	_level int64
}

// NewAlitripplatformdivisionsgetdivisionbynameRequest 初始化AlitripplatformdivisionsgetdivisionbynameAPIRequest对象
func NewAlitripplatformdivisionsgetdivisionbynameRequest() *AlitripplatformdivisionsgetdivisionbynameAPIRequest {
	return &AlitripplatformdivisionsgetdivisionbynameAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripplatformdivisionsgetdivisionbynameAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.divisions.getdivisionbyname"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripplatformdivisionsgetdivisionbynameAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripplatformdivisionsgetdivisionbynameAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 行政区划名称
func (r *AlitripplatformdivisionsgetdivisionbynameAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlitripplatformdivisionsgetdivisionbynameAPIRequest) GetName() string {
	return r._name
}

// SetLevel is Level Setter
// 行政区划级别ALL(0, &#34;全部&#34;),  	CONTINENT(1, &#34;大洲&#34;),  	COUNTRY(2, &#34;国家&#34;),  	PROVINCE(3, &#34;省份&#34;),  	CITY(4, &#34;城市&#34;),  	DISTRICT(5, &#34;区县&#34;),  	STREET(6, &#34;街道&#34;)
func (r *AlitripplatformdivisionsgetdivisionbynameAPIRequest) SetLevel(_level int64) error {
	r._level = _level
	r.Set("level", _level)
	return nil
}

// GetLevel Level Getter
func (r AlitripplatformdivisionsgetdivisionbynameAPIRequest) GetLevel() int64 {
	return r._level
}
