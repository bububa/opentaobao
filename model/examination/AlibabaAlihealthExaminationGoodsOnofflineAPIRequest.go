package examination

import (
	"net/url"
	"sync"

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
	// 门店code列表
	_hospitalCodes string
	// 操作类型: online=上线，offline=下线
	_type string
}

// NewAlibabaAlihealthExaminationGoodsOnofflineRequest 初始化AlibabaAlihealthExaminationGoodsOnofflineAPIRequest对象
func NewAlibabaAlihealthExaminationGoodsOnofflineRequest() *AlibabaAlihealthExaminationGoodsOnofflineAPIRequest {
	return &AlibabaAlihealthExaminationGoodsOnofflineAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthExaminationGoodsOnofflineAPIRequest) Reset() {
	r._groupId = ""
	r._hospitalCodes = ""
	r._type = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationGoodsOnofflineAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.goods.onoffline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationGoodsOnofflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthExaminationGoodsOnofflineAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihealthExaminationGoodsOnofflineAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthExaminationGoodsOnofflineRequest()
	},
}

// GetAlibabaAlihealthExaminationGoodsOnofflineRequest 从 sync.Pool 获取 AlibabaAlihealthExaminationGoodsOnofflineAPIRequest
func GetAlibabaAlihealthExaminationGoodsOnofflineAPIRequest() *AlibabaAlihealthExaminationGoodsOnofflineAPIRequest {
	return poolAlibabaAlihealthExaminationGoodsOnofflineAPIRequest.Get().(*AlibabaAlihealthExaminationGoodsOnofflineAPIRequest)
}

// ReleaseAlibabaAlihealthExaminationGoodsOnofflineAPIRequest 将 AlibabaAlihealthExaminationGoodsOnofflineAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthExaminationGoodsOnofflineAPIRequest(v *AlibabaAlihealthExaminationGoodsOnofflineAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthExaminationGoodsOnofflineAPIRequest.Put(v)
}
