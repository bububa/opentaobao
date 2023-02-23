package damai

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
