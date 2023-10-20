package normalvisa

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelVisaSignSendAPIRequest 签证批量申请人送签接口 API请求
// alitrip.travel.visa.sign.send
//
// 签证批量申请人送签接口，用于商家批量送签。
type AlitripTravelVisaSignSendAPIRequest struct {
	model.Params
	// 申请人ids
	_applyIds []string
	// 国家id。目前只支持越南，越南国家id:27027
	_nationId int64
	// 送签类型：1-非加急，2-加急，默认非加急
	_signType int64
}

// NewAlitripTravelVisaSignSendRequest 初始化AlitripTravelVisaSignSendAPIRequest对象
func NewAlitripTravelVisaSignSendRequest() *AlitripTravelVisaSignSendAPIRequest {
	return &AlitripTravelVisaSignSendAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTravelVisaSignSendAPIRequest) Reset() {
	r._applyIds = r._applyIds[:0]
	r._nationId = 0
	r._signType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelVisaSignSendAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.visa.sign.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelVisaSignSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTravelVisaSignSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyIds is ApplyIds Setter
// 申请人ids
func (r *AlitripTravelVisaSignSendAPIRequest) SetApplyIds(_applyIds []string) error {
	r._applyIds = _applyIds
	r.Set("apply_ids", _applyIds)
	return nil
}

// GetApplyIds ApplyIds Getter
func (r AlitripTravelVisaSignSendAPIRequest) GetApplyIds() []string {
	return r._applyIds
}

// SetNationId is NationId Setter
// 国家id。目前只支持越南，越南国家id:27027
func (r *AlitripTravelVisaSignSendAPIRequest) SetNationId(_nationId int64) error {
	r._nationId = _nationId
	r.Set("nation_id", _nationId)
	return nil
}

// GetNationId NationId Getter
func (r AlitripTravelVisaSignSendAPIRequest) GetNationId() int64 {
	return r._nationId
}

// SetSignType is SignType Setter
// 送签类型：1-非加急，2-加急，默认非加急
func (r *AlitripTravelVisaSignSendAPIRequest) SetSignType(_signType int64) error {
	r._signType = _signType
	r.Set("sign_type", _signType)
	return nil
}

// GetSignType SignType Getter
func (r AlitripTravelVisaSignSendAPIRequest) GetSignType() int64 {
	return r._signType
}

var poolAlitripTravelVisaSignSendAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTravelVisaSignSendRequest()
	},
}

// GetAlitripTravelVisaSignSendRequest 从 sync.Pool 获取 AlitripTravelVisaSignSendAPIRequest
func GetAlitripTravelVisaSignSendAPIRequest() *AlitripTravelVisaSignSendAPIRequest {
	return poolAlitripTravelVisaSignSendAPIRequest.Get().(*AlitripTravelVisaSignSendAPIRequest)
}

// ReleaseAlitripTravelVisaSignSendAPIRequest 将 AlitripTravelVisaSignSendAPIRequest 放入 sync.Pool
func ReleaseAlitripTravelVisaSignSendAPIRequest(v *AlitripTravelVisaSignSendAPIRequest) {
	v.Reset()
	poolAlitripTravelVisaSignSendAPIRequest.Put(v)
}
