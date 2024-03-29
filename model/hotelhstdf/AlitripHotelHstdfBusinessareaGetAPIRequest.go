package hotelhstdf

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfBusinessareaGetAPIRequest 根据城市查询商圈 API请求
// alitrip.hotel.hstdf.businessarea.get
//
// 根据cityId分页查询商圈信息
type AlitripHotelHstdfBusinessareaGetAPIRequest struct {
	model.Params
	// 请求参数封装
	_paramGetByTrdiDivisionIdParam *GetByTrdiDivisionIdParam
}

// NewAlitripHotelHstdfBusinessareaGetRequest 初始化AlitripHotelHstdfBusinessareaGetAPIRequest对象
func NewAlitripHotelHstdfBusinessareaGetRequest() *AlitripHotelHstdfBusinessareaGetAPIRequest {
	return &AlitripHotelHstdfBusinessareaGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripHotelHstdfBusinessareaGetAPIRequest) Reset() {
	r._paramGetByTrdiDivisionIdParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfBusinessareaGetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.businessarea.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfBusinessareaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripHotelHstdfBusinessareaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamGetByTrdiDivisionIdParam is ParamGetByTrdiDivisionIdParam Setter
// 请求参数封装
func (r *AlitripHotelHstdfBusinessareaGetAPIRequest) SetParamGetByTrdiDivisionIdParam(_paramGetByTrdiDivisionIdParam *GetByTrdiDivisionIdParam) error {
	r._paramGetByTrdiDivisionIdParam = _paramGetByTrdiDivisionIdParam
	r.Set("param_get_by_trdi_division_id_param", _paramGetByTrdiDivisionIdParam)
	return nil
}

// GetParamGetByTrdiDivisionIdParam ParamGetByTrdiDivisionIdParam Getter
func (r AlitripHotelHstdfBusinessareaGetAPIRequest) GetParamGetByTrdiDivisionIdParam() *GetByTrdiDivisionIdParam {
	return r._paramGetByTrdiDivisionIdParam
}

var poolAlitripHotelHstdfBusinessareaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripHotelHstdfBusinessareaGetRequest()
	},
}

// GetAlitripHotelHstdfBusinessareaGetRequest 从 sync.Pool 获取 AlitripHotelHstdfBusinessareaGetAPIRequest
func GetAlitripHotelHstdfBusinessareaGetAPIRequest() *AlitripHotelHstdfBusinessareaGetAPIRequest {
	return poolAlitripHotelHstdfBusinessareaGetAPIRequest.Get().(*AlitripHotelHstdfBusinessareaGetAPIRequest)
}

// ReleaseAlitripHotelHstdfBusinessareaGetAPIRequest 将 AlitripHotelHstdfBusinessareaGetAPIRequest 放入 sync.Pool
func ReleaseAlitripHotelHstdfBusinessareaGetAPIRequest(v *AlitripHotelHstdfBusinessareaGetAPIRequest) {
	v.Reset()
	poolAlitripHotelHstdfBusinessareaGetAPIRequest.Put(v)
}
