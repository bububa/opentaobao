package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrugcodesAPIRequest 药品是否赋码 API请求
// alibaba.alihealth.drug.kyt.drugcodes
//
// 药品是否赋码
type AlibabaAlihealthDrugKytDrugcodesAPIRequest struct {
	model.Params
	// 企业名称
	_refEntName string
	// 药品名称
	_physicName string
	// 生产批号
	_produceBatchNo string
	// 药品类型
	_physicType string
	// 包装规格
	_pkgSpec string
	// 制剂规格
	_prepnSpec string
}

// NewAlibabaAlihealthDrugKytDrugcodesRequest 初始化AlibabaAlihealthDrugKytDrugcodesAPIRequest对象
func NewAlibabaAlihealthDrugKytDrugcodesRequest() *AlibabaAlihealthDrugKytDrugcodesAPIRequest {
	return &AlibabaAlihealthDrugKytDrugcodesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrugcodesAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.drugcodes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrugcodesAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytDrugcodesAPIRequest) SetRefEntName(_refEntName string) error {
	r._refEntName = _refEntName
	r.Set("ref_ent_name", _refEntName)
	return nil
}

// Get RefEntName Getter
func (r AlibabaAlihealthDrugKytDrugcodesAPIRequest) GetRefEntName() string {
	return r._refEntName
}

// Set is PhysicName Setter
// 药品名称
func (r *AlibabaAlihealthDrugKytDrugcodesAPIRequest) SetPhysicName(_physicName string) error {
	r._physicName = _physicName
	r.Set("physic_name", _physicName)
	return nil
}

// Get PhysicName Getter
func (r AlibabaAlihealthDrugKytDrugcodesAPIRequest) GetPhysicName() string {
	return r._physicName
}

// Set is ProduceBatchNo Setter
// 生产批号
func (r *AlibabaAlihealthDrugKytDrugcodesAPIRequest) SetProduceBatchNo(_produceBatchNo string) error {
	r._produceBatchNo = _produceBatchNo
	r.Set("produce_batch_no", _produceBatchNo)
	return nil
}

// Get ProduceBatchNo Getter
func (r AlibabaAlihealthDrugKytDrugcodesAPIRequest) GetProduceBatchNo() string {
	return r._produceBatchNo
}

// Set is PhysicType Setter
// 药品类型
func (r *AlibabaAlihealthDrugKytDrugcodesAPIRequest) SetPhysicType(_physicType string) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// Get PhysicType Getter
func (r AlibabaAlihealthDrugKytDrugcodesAPIRequest) GetPhysicType() string {
	return r._physicType
}

// Set is PkgSpec Setter
// 包装规格
func (r *AlibabaAlihealthDrugKytDrugcodesAPIRequest) SetPkgSpec(_pkgSpec string) error {
	r._pkgSpec = _pkgSpec
	r.Set("pkg_spec", _pkgSpec)
	return nil
}

// Get PkgSpec Getter
func (r AlibabaAlihealthDrugKytDrugcodesAPIRequest) GetPkgSpec() string {
	return r._pkgSpec
}

// Set is PrepnSpec Setter
// 制剂规格
func (r *AlibabaAlihealthDrugKytDrugcodesAPIRequest) SetPrepnSpec(_prepnSpec string) error {
	r._prepnSpec = _prepnSpec
	r.Set("prepn_spec", _prepnSpec)
	return nil
}

// Get PrepnSpec Getter
func (r AlibabaAlihealthDrugKytDrugcodesAPIRequest) GetPrepnSpec() string {
	return r._prepnSpec
}
