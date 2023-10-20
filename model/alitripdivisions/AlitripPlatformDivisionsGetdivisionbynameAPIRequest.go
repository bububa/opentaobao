package alitripdivisions

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPlatformDivisionsGetdivisionbynameAPIRequest 根据中文名称与行政区划级别查询行政区划数据 API请求
// alitrip.platform.divisions.getdivisionbyname
//
// 根据中文名称与行政区划级别查询行政区划数据
type AlitripPlatformDivisionsGetdivisionbynameAPIRequest struct {
	model.Params
	// 行政区划名称
	_name string
	// 行政区划级别ALL(0, "全部"),  	CONTINENT(1, "大洲"),  	COUNTRY(2, "国家"),  	PROVINCE(3, "省份"),  	CITY(4, "城市"),  	DISTRICT(5, "区县"),  	STREET(6, "街道")
	_level int64
}

// NewAlitripPlatformDivisionsGetdivisionbynameRequest 初始化AlitripPlatformDivisionsGetdivisionbynameAPIRequest对象
func NewAlitripPlatformDivisionsGetdivisionbynameRequest() *AlitripPlatformDivisionsGetdivisionbynameAPIRequest {
	return &AlitripPlatformDivisionsGetdivisionbynameAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripPlatformDivisionsGetdivisionbynameAPIRequest) Reset() {
	r._name = ""
	r._level = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPlatformDivisionsGetdivisionbynameAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.divisions.getdivisionbyname"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPlatformDivisionsGetdivisionbynameAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripPlatformDivisionsGetdivisionbynameAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 行政区划名称
func (r *AlitripPlatformDivisionsGetdivisionbynameAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlitripPlatformDivisionsGetdivisionbynameAPIRequest) GetName() string {
	return r._name
}

// SetLevel is Level Setter
// 行政区划级别ALL(0, &#34;全部&#34;),  	CONTINENT(1, &#34;大洲&#34;),  	COUNTRY(2, &#34;国家&#34;),  	PROVINCE(3, &#34;省份&#34;),  	CITY(4, &#34;城市&#34;),  	DISTRICT(5, &#34;区县&#34;),  	STREET(6, &#34;街道&#34;)
func (r *AlitripPlatformDivisionsGetdivisionbynameAPIRequest) SetLevel(_level int64) error {
	r._level = _level
	r.Set("level", _level)
	return nil
}

// GetLevel Level Getter
func (r AlitripPlatformDivisionsGetdivisionbynameAPIRequest) GetLevel() int64 {
	return r._level
}

var poolAlitripPlatformDivisionsGetdivisionbynameAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripPlatformDivisionsGetdivisionbynameRequest()
	},
}

// GetAlitripPlatformDivisionsGetdivisionbynameRequest 从 sync.Pool 获取 AlitripPlatformDivisionsGetdivisionbynameAPIRequest
func GetAlitripPlatformDivisionsGetdivisionbynameAPIRequest() *AlitripPlatformDivisionsGetdivisionbynameAPIRequest {
	return poolAlitripPlatformDivisionsGetdivisionbynameAPIRequest.Get().(*AlitripPlatformDivisionsGetdivisionbynameAPIRequest)
}

// ReleaseAlitripPlatformDivisionsGetdivisionbynameAPIRequest 将 AlitripPlatformDivisionsGetdivisionbynameAPIRequest 放入 sync.Pool
func ReleaseAlitripPlatformDivisionsGetdivisionbynameAPIRequest(v *AlitripPlatformDivisionsGetdivisionbynameAPIRequest) {
	v.Reset()
	poolAlitripPlatformDivisionsGetdivisionbynameAPIRequest.Put(v)
}
