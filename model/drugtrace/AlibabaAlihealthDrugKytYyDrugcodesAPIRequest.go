package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytyydrugcodesAPIRequest 查询药品是否赋码 API请求
// alibaba.alihealth.drug.kyt.yy.drugcodes
//
// 药品是否赋码
type AlibabaalihealthdrugkytyydrugcodesAPIRequest struct {
	model.Params
	// 企业名称
	_refEntName string
	// 药品名称
	_physicName string
	// 生产批号
	_produceBatchNo string
	// 药品类型
	_physicType string
	// 制剂规格
	_prepnSpec string
	// 包装规格
	_pkgSpec string
}

// NewAlibabaalihealthdrugkytyydrugcodesRequest 初始化AlibabaalihealthdrugkytyydrugcodesAPIRequest对象
func NewAlibabaalihealthdrugkytyydrugcodesRequest() *AlibabaalihealthdrugkytyydrugcodesAPIRequest {
	return &AlibabaalihealthdrugkytyydrugcodesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytyydrugcodesAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.yy.drugcodes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytyydrugcodesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytyydrugcodesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntName is RefEntName Setter
// 企业名称
func (r *AlibabaalihealthdrugkytyydrugcodesAPIRequest) SetRefEntName(_refEntName string) error {
	r._refEntName = _refEntName
	r.Set("ref_ent_name", _refEntName)
	return nil
}

// GetRefEntName RefEntName Getter
func (r AlibabaalihealthdrugkytyydrugcodesAPIRequest) GetRefEntName() string {
	return r._refEntName
}

// SetPhysicName is PhysicName Setter
// 药品名称
func (r *AlibabaalihealthdrugkytyydrugcodesAPIRequest) SetPhysicName(_physicName string) error {
	r._physicName = _physicName
	r.Set("physic_name", _physicName)
	return nil
}

// GetPhysicName PhysicName Getter
func (r AlibabaalihealthdrugkytyydrugcodesAPIRequest) GetPhysicName() string {
	return r._physicName
}

// SetProduceBatchNo is ProduceBatchNo Setter
// 生产批号
func (r *AlibabaalihealthdrugkytyydrugcodesAPIRequest) SetProduceBatchNo(_produceBatchNo string) error {
	r._produceBatchNo = _produceBatchNo
	r.Set("produce_batch_no", _produceBatchNo)
	return nil
}

// GetProduceBatchNo ProduceBatchNo Getter
func (r AlibabaalihealthdrugkytyydrugcodesAPIRequest) GetProduceBatchNo() string {
	return r._produceBatchNo
}

// SetPhysicType is PhysicType Setter
// 药品类型
func (r *AlibabaalihealthdrugkytyydrugcodesAPIRequest) SetPhysicType(_physicType string) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaalihealthdrugkytyydrugcodesAPIRequest) GetPhysicType() string {
	return r._physicType
}

// SetPrepnSpec is PrepnSpec Setter
// 制剂规格
func (r *AlibabaalihealthdrugkytyydrugcodesAPIRequest) SetPrepnSpec(_prepnSpec string) error {
	r._prepnSpec = _prepnSpec
	r.Set("prepn_spec", _prepnSpec)
	return nil
}

// GetPrepnSpec PrepnSpec Getter
func (r AlibabaalihealthdrugkytyydrugcodesAPIRequest) GetPrepnSpec() string {
	return r._prepnSpec
}

// SetPkgSpec is PkgSpec Setter
// 包装规格
func (r *AlibabaalihealthdrugkytyydrugcodesAPIRequest) SetPkgSpec(_pkgSpec string) error {
	r._pkgSpec = _pkgSpec
	r.Set("pkg_spec", _pkgSpec)
	return nil
}

// GetPkgSpec PkgSpec Getter
func (r AlibabaalihealthdrugkytyydrugcodesAPIRequest) GetPkgSpec() string {
	return r._pkgSpec
}
