package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelVisaSignSendAPIRequest 签证批量申请人送签接口 API请求
// alitrip.travel.visa.sign.send
//
// 签证批量申请人送签接口，用于商家批量送签。
type AlitripTravelVisaSignSendAPIRequest struct {
	model.Params
	// 国家id。目前只支持越南，越南国家id:27027
	_nationId int64
	// 送签类型：1-非加急，2-加急，默认非加急
	_signType int64
	// 申请人ids
	_applyIds []string
}

// NewAlitripTravelVisaSignSendRequest 初始化AlitripTravelVisaSignSendAPIRequest对象
func NewAlitripTravelVisaSignSendRequest() *AlitripTravelVisaSignSendAPIRequest {
	return &AlitripTravelVisaSignSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelVisaSignSendAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.visa.sign.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelVisaSignSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is NationId Setter
// 国家id。目前只支持越南，越南国家id:27027
func (r *AlitripTravelVisaSignSendAPIRequest) SetNationId(_nationId int64) error {
	r._nationId = _nationId
	r.Set("nation_id", _nationId)
	return nil
}

// Get NationId Getter
func (r AlitripTravelVisaSignSendAPIRequest) GetNationId() int64 {
	return r._nationId
}

// Set is SignType Setter
// 送签类型：1-非加急，2-加急，默认非加急
func (r *AlitripTravelVisaSignSendAPIRequest) SetSignType(_signType int64) error {
	r._signType = _signType
	r.Set("sign_type", _signType)
	return nil
}

// Get SignType Getter
func (r AlitripTravelVisaSignSendAPIRequest) GetSignType() int64 {
	return r._signType
}

// Set is ApplyIds Setter
// 申请人ids
func (r *AlitripTravelVisaSignSendAPIRequest) SetApplyIds(_applyIds []string) error {
	r._applyIds = _applyIds
	r.Set("apply_ids", _applyIds)
	return nil
}

// Get ApplyIds Getter
func (r AlitripTravelVisaSignSendAPIRequest) GetApplyIds() []string {
	return r._applyIds
}
