package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSinglerelationAPIRequest 单码关联关系查询，通过一个码查询这个码下的所有子码 API请求
// alibaba.alihealth.drug.kyt.singlerelation
//
// 单码关联关系查询，通过一个码查询这个码下的所有子码。（只有做过入库的码，才能能进行查询）
type AlibabaAlihealthDrugKytSinglerelationAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
	_desRefEntId string
}

// NewAlibabaAlihealthDrugKytSinglerelationRequest 初始化AlibabaAlihealthDrugKytSinglerelationAPIRequest对象
func NewAlibabaAlihealthDrugKytSinglerelationRequest() *AlibabaAlihealthDrugKytSinglerelationAPIRequest {
	return &AlibabaAlihealthDrugKytSinglerelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSinglerelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.singlerelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSinglerelationAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytSinglerelationAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugKytSinglerelationAPIRequest) GetCode() string {
	return r._code
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytSinglerelationAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytSinglerelationAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetDesRefEntId is DesRefEntId Setter
// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
func (r *AlibabaAlihealthDrugKytSinglerelationAPIRequest) SetDesRefEntId(_desRefEntId string) error {
	r._desRefEntId = _desRefEntId
	r.Set("des_ref_ent_id", _desRefEntId)
	return nil
}

// GetDesRefEntId DesRefEntId Getter
func (r AlibabaAlihealthDrugKytSinglerelationAPIRequest) GetDesRefEntId() string {
	return r._desRefEntId
}
