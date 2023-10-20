package hotelhstdf

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfPoilocationGetAPIRequest 根据平台城市id分页查询poi location API请求
// alitrip.hotel.hstdf.poilocation.get
//
// 根据平台城市id分页查询poi location
type AlitripHotelHstdfPoilocationGetAPIRequest struct {
	model.Params
	// 参数封装
	_paramGetByTrdiDivisionIdParam *GetByTrdiDivisionIdParam
}

// NewAlitripHotelHstdfPoilocationGetRequest 初始化AlitripHotelHstdfPoilocationGetAPIRequest对象
func NewAlitripHotelHstdfPoilocationGetRequest() *AlitripHotelHstdfPoilocationGetAPIRequest {
	return &AlitripHotelHstdfPoilocationGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripHotelHstdfPoilocationGetAPIRequest) Reset() {
	r._paramGetByTrdiDivisionIdParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfPoilocationGetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.poilocation.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfPoilocationGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripHotelHstdfPoilocationGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamGetByTrdiDivisionIdParam is ParamGetByTrdiDivisionIdParam Setter
// 参数封装
func (r *AlitripHotelHstdfPoilocationGetAPIRequest) SetParamGetByTrdiDivisionIdParam(_paramGetByTrdiDivisionIdParam *GetByTrdiDivisionIdParam) error {
	r._paramGetByTrdiDivisionIdParam = _paramGetByTrdiDivisionIdParam
	r.Set("param_get_by_trdi_division_id_param", _paramGetByTrdiDivisionIdParam)
	return nil
}

// GetParamGetByTrdiDivisionIdParam ParamGetByTrdiDivisionIdParam Getter
func (r AlitripHotelHstdfPoilocationGetAPIRequest) GetParamGetByTrdiDivisionIdParam() *GetByTrdiDivisionIdParam {
	return r._paramGetByTrdiDivisionIdParam
}

var poolAlitripHotelHstdfPoilocationGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripHotelHstdfPoilocationGetRequest()
	},
}

// GetAlitripHotelHstdfPoilocationGetRequest 从 sync.Pool 获取 AlitripHotelHstdfPoilocationGetAPIRequest
func GetAlitripHotelHstdfPoilocationGetAPIRequest() *AlitripHotelHstdfPoilocationGetAPIRequest {
	return poolAlitripHotelHstdfPoilocationGetAPIRequest.Get().(*AlitripHotelHstdfPoilocationGetAPIRequest)
}

// ReleaseAlitripHotelHstdfPoilocationGetAPIRequest 将 AlitripHotelHstdfPoilocationGetAPIRequest 放入 sync.Pool
func ReleaseAlitripHotelHstdfPoilocationGetAPIRequest(v *AlitripHotelHstdfPoilocationGetAPIRequest) {
	v.Reset()
	poolAlitripHotelHstdfPoilocationGetAPIRequest.Put(v)
}
