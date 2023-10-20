package customizemarket

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTaobaoIndustryPetProfileQueryAPIRequest 用户宠物列表查询 API请求
// alibaba.taobao.industry.pet.profile.query
//
// 用户宠物列表查询
type AlibabaTaobaoIndustryPetProfileQueryAPIRequest struct {
	model.Params
	// 参数
	_profileQuery *ProfileQuery
}

// NewAlibabaTaobaoIndustryPetProfileQueryRequest 初始化AlibabaTaobaoIndustryPetProfileQueryAPIRequest对象
func NewAlibabaTaobaoIndustryPetProfileQueryRequest() *AlibabaTaobaoIndustryPetProfileQueryAPIRequest {
	return &AlibabaTaobaoIndustryPetProfileQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTaobaoIndustryPetProfileQueryAPIRequest) Reset() {
	r._profileQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTaobaoIndustryPetProfileQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.taobao.industry.pet.profile.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTaobaoIndustryPetProfileQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTaobaoIndustryPetProfileQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProfileQuery is ProfileQuery Setter
// 参数
func (r *AlibabaTaobaoIndustryPetProfileQueryAPIRequest) SetProfileQuery(_profileQuery *ProfileQuery) error {
	r._profileQuery = _profileQuery
	r.Set("profile_query", _profileQuery)
	return nil
}

// GetProfileQuery ProfileQuery Getter
func (r AlibabaTaobaoIndustryPetProfileQueryAPIRequest) GetProfileQuery() *ProfileQuery {
	return r._profileQuery
}

var poolAlibabaTaobaoIndustryPetProfileQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTaobaoIndustryPetProfileQueryRequest()
	},
}

// GetAlibabaTaobaoIndustryPetProfileQueryRequest 从 sync.Pool 获取 AlibabaTaobaoIndustryPetProfileQueryAPIRequest
func GetAlibabaTaobaoIndustryPetProfileQueryAPIRequest() *AlibabaTaobaoIndustryPetProfileQueryAPIRequest {
	return poolAlibabaTaobaoIndustryPetProfileQueryAPIRequest.Get().(*AlibabaTaobaoIndustryPetProfileQueryAPIRequest)
}

// ReleaseAlibabaTaobaoIndustryPetProfileQueryAPIRequest 将 AlibabaTaobaoIndustryPetProfileQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaTaobaoIndustryPetProfileQueryAPIRequest(v *AlibabaTaobaoIndustryPetProfileQueryAPIRequest) {
	v.Reset()
	poolAlibabaTaobaoIndustryPetProfileQueryAPIRequest.Put(v)
}
