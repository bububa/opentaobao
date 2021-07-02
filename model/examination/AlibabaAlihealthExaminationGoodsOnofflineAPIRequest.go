package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationGoodsOnofflineAPIRequest 上线/下线 体检产品 API请求
// alibaba.alihealth.examination.goods.onoffline
//
// 第三方体检机构对接钉钉体检中的产品 上线／下线
type AlibabaAlihealthExaminationGoodsOnofflineAPIRequest struct {
	model.Params
	// 商品组code，机构保证唯一
	_groupId string
	// 操作类型: online=上线，offline=下线
	_type string
	// 门店code列表
	_hospitalCodes string
}

// NewAlibabaAlihealthExaminationGoodsOnofflineRequest 初始化AlibabaAlihealthExaminationGoodsOnofflineAPIRequest对象
func NewAlibabaAlihealthExaminationGoodsOnofflineRequest() *AlibabaAlihealthExaminationGoodsOnofflineAPIRequest {
	return &AlibabaAlihealthExaminationGoodsOnofflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationGoodsOnofflineAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.goods.onoffline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationGoodsOnofflineAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetGroupId is GroupId Setter
// 商品组code，机构保证唯一
func (r *AlibabaAlihealthExaminationGoodsOnofflineAPIRequest) SetGroupId(_groupId string) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabaAlihealthExaminationGoodsOnofflineAPIRequest) GetGroupId() string {
	return r._groupId
}

// SetType is Type Setter
// 操作类型: online=上线，offline=下线
func (r *AlibabaAlihealthExaminationGoodsOnofflineAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaAlihealthExaminationGoodsOnofflineAPIRequest) GetType() string {
	return r._type
}

// SetHospitalCodes is HospitalCodes Setter
// 门店code列表
func (r *AlibabaAlihealthExaminationGoodsOnofflineAPIRequest) SetHospitalCodes(_hospitalCodes string) error {
	r._hospitalCodes = _hospitalCodes
	r.Set("hospital_codes", _hospitalCodes)
	return nil
}

// GetHospitalCodes HospitalCodes Getter
func (r AlibabaAlihealthExaminationGoodsOnofflineAPIRequest) GetHospitalCodes() string {
	return r._hospitalCodes
}
