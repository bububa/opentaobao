package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationgoodsonofflineAPIRequest 上线/下线 体检产品 API请求
// alibaba.alihealth.examination.goods.onoffline
//
// 第三方体检机构对接钉钉体检中的产品 上线／下线
type AlibabaalihealthexaminationgoodsonofflineAPIRequest struct {
	model.Params
	// 商品组code，机构保证唯一
	_groupId string
	// 门店code列表
	_hospitalCodes string
	// 操作类型: online=上线，offline=下线
	_type string
}

// NewAlibabaalihealthexaminationgoodsonofflineRequest 初始化AlibabaalihealthexaminationgoodsonofflineAPIRequest对象
func NewAlibabaalihealthexaminationgoodsonofflineRequest() *AlibabaalihealthexaminationgoodsonofflineAPIRequest {
	return &AlibabaalihealthexaminationgoodsonofflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationgoodsonofflineAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.goods.onoffline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationgoodsonofflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationgoodsonofflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupId is GroupId Setter
// 商品组code，机构保证唯一
func (r *AlibabaalihealthexaminationgoodsonofflineAPIRequest) SetGroupId(_groupId string) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabaalihealthexaminationgoodsonofflineAPIRequest) GetGroupId() string {
	return r._groupId
}

// SetHospitalCodes is HospitalCodes Setter
// 门店code列表
func (r *AlibabaalihealthexaminationgoodsonofflineAPIRequest) SetHospitalCodes(_hospitalCodes string) error {
	r._hospitalCodes = _hospitalCodes
	r.Set("hospital_codes", _hospitalCodes)
	return nil
}

// GetHospitalCodes HospitalCodes Getter
func (r AlibabaalihealthexaminationgoodsonofflineAPIRequest) GetHospitalCodes() string {
	return r._hospitalCodes
}

// SetType is Type Setter
// 操作类型: online=上线，offline=下线
func (r *AlibabaalihealthexaminationgoodsonofflineAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaalihealthexaminationgoodsonofflineAPIRequest) GetType() string {
	return r._type
}
