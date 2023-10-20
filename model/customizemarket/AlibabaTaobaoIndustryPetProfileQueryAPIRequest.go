package customizemarket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabataobaoindustrypetprofilequeryAPIRequest 用户宠物列表查询 API请求
// alibaba.taobao.industry.pet.profile.query
//
// 用户宠物列表查询
type AlibabataobaoindustrypetprofilequeryAPIRequest struct {
	model.Params
	// 参数
	_profileQuery *ProfileQuery
}

// NewAlibabataobaoindustrypetprofilequeryRequest 初始化AlibabataobaoindustrypetprofilequeryAPIRequest对象
func NewAlibabataobaoindustrypetprofilequeryRequest() *AlibabataobaoindustrypetprofilequeryAPIRequest {
	return &AlibabataobaoindustrypetprofilequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabataobaoindustrypetprofilequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.taobao.industry.pet.profile.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabataobaoindustrypetprofilequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabataobaoindustrypetprofilequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProfileQuery is ProfileQuery Setter
// 参数
func (r *AlibabataobaoindustrypetprofilequeryAPIRequest) SetProfileQuery(_profileQuery *ProfileQuery) error {
	r._profileQuery = _profileQuery
	r.Set("profile_query", _profileQuery)
	return nil
}

// GetProfileQuery ProfileQuery Getter
func (r AlibabataobaoindustrypetprofilequeryAPIRequest) GetProfileQuery() *ProfileQuery {
	return r._profileQuery
}
