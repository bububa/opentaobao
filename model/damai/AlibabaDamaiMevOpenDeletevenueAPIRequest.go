package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeletevenueAPIRequest 大麦换验平台-第三方对外开放-场馆接口deleteVenue API请求
// alibaba.damai.mev.open.deletevenue
//
// 开放接口，删除场馆
type AlibabaDamaiMevOpenDeletevenueAPIRequest struct {
	model.Params
	// 入参deleteVenueParam
	_deleteVenueParam *VenueIdOpenParam
}

// NewAlibabaDamaiMevOpenDeletevenueRequest 初始化AlibabaDamaiMevOpenDeletevenueAPIRequest对象
func NewAlibabaDamaiMevOpenDeletevenueRequest() *AlibabaDamaiMevOpenDeletevenueAPIRequest {
	return &AlibabaDamaiMevOpenDeletevenueAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenDeletevenueAPIRequest) Reset() {
	r._deleteVenueParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeletevenueAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.deletevenue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeletevenueAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenDeletevenueAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteVenueParam is DeleteVenueParam Setter
// 入参deleteVenueParam
func (r *AlibabaDamaiMevOpenDeletevenueAPIRequest) SetDeleteVenueParam(_deleteVenueParam *VenueIdOpenParam) error {
	r._deleteVenueParam = _deleteVenueParam
	r.Set("delete_venue_param", _deleteVenueParam)
	return nil
}

// GetDeleteVenueParam DeleteVenueParam Getter
func (r AlibabaDamaiMevOpenDeletevenueAPIRequest) GetDeleteVenueParam() *VenueIdOpenParam {
	return r._deleteVenueParam
}

var poolAlibabaDamaiMevOpenDeletevenueAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenDeletevenueRequest()
	},
}

// GetAlibabaDamaiMevOpenDeletevenueRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenDeletevenueAPIRequest
func GetAlibabaDamaiMevOpenDeletevenueAPIRequest() *AlibabaDamaiMevOpenDeletevenueAPIRequest {
	return poolAlibabaDamaiMevOpenDeletevenueAPIRequest.Get().(*AlibabaDamaiMevOpenDeletevenueAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenDeletevenueAPIRequest 将 AlibabaDamaiMevOpenDeletevenueAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenDeletevenueAPIRequest(v *AlibabaDamaiMevOpenDeletevenueAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenDeletevenueAPIRequest.Put(v)
}
