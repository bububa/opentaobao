package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytYyDrugcodesAPIRequest 查询药品是否赋码 API请求
// alibaba.alihealth.drug.kyt.yy.drugcodes
//
// 药品是否赋码
type AlibabaAlihealthDrugKytYyDrugcodesAPIRequest struct {
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

// NewAlibabaAlihealthDrugKytYyDrugcodesRequest 初始化AlibabaAlihealthDrugKytYyDrugcodesAPIRequest对象
func NewAlibabaAlihealthDrugKytYyDrugcodesRequest() *AlibabaAlihealthDrugKytYyDrugcodesAPIRequest {
	return &AlibabaAlihealthDrugKytYyDrugcodesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytYyDrugcodesAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.yy.drugcodes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytYyDrugcodesAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRefEntName is RefEntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytYyDrugcodesAPIRequest) SetRefEntName(_refEntName string) error {
	r._refEntName = _refEntName
	r.Set("ref_ent_name", _refEntName)
	return nil
}

// GetRefEntName RefEntName Getter
func (r AlibabaAlihealthDrugKytYyDrugcodesAPIRequest) GetRefEntName() string {
	return r._refEntName
}

// SetPhysicName is PhysicName Setter
// 药品名称
func (r *AlibabaAlihealthDrugKytYyDrugcodesAPIRequest) SetPhysicName(_physicName string) error {
	r._physicName = _physicName
	r.Set("physic_name", _physicName)
	return nil
}

// GetPhysicName PhysicName Getter
func (r AlibabaAlihealthDrugKytYyDrugcodesAPIRequest) GetPhysicName() string {
	return r._physicName
}

// SetProduceBatchNo is ProduceBatchNo Setter
// 生产批号
func (r *AlibabaAlihealthDrugKytYyDrugcodesAPIRequest) SetProduceBatchNo(_produceBatchNo string) error {
	r._produceBatchNo = _produceBatchNo
	r.Set("produce_batch_no", _produceBatchNo)
	return nil
}

// GetProduceBatchNo ProduceBatchNo Getter
func (r AlibabaAlihealthDrugKytYyDrugcodesAPIRequest) GetProduceBatchNo() string {
	return r._produceBatchNo
}

// SetPhysicType is PhysicType Setter
// 药品类型
func (r *AlibabaAlihealthDrugKytYyDrugcodesAPIRequest) SetPhysicType(_physicType string) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaAlihealthDrugKytYyDrugcodesAPIRequest) GetPhysicType() string {
	return r._physicType
}

// SetPkgSpec is PkgSpec Setter
// 包装规格
func (r *AlibabaAlihealthDrugKytYyDrugcodesAPIRequest) SetPkgSpec(_pkgSpec string) error {
	r._pkgSpec = _pkgSpec
	r.Set("pkg_spec", _pkgSpec)
	return nil
}

// GetPkgSpec PkgSpec Getter
func (r AlibabaAlihealthDrugKytYyDrugcodesAPIRequest) GetPkgSpec() string {
	return r._pkgSpec
}

// SetPrepnSpec is PrepnSpec Setter
// 制剂规格
func (r *AlibabaAlihealthDrugKytYyDrugcodesAPIRequest) SetPrepnSpec(_prepnSpec string) error {
	r._prepnSpec = _prepnSpec
	r.Set("prepn_spec", _prepnSpec)
	return nil
}

// GetPrepnSpec PrepnSpec Getter
func (r AlibabaAlihealthDrugKytYyDrugcodesAPIRequest) GetPrepnSpec() string {
	return r._prepnSpec
}
