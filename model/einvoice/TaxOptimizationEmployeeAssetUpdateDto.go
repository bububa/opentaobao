package einvoice

import (
	"sync"
)

// TaxOptimizationEmployeeAssetUpdateDto 结构体
type TaxOptimizationEmployeeAssetUpdateDto struct {
	// 需要更新的资产账号
	AssetSymbol string `json:"asset_symbol,omitempty" xml:"asset_symbol,omitempty"`
	// 需要更新的资产类型
	AssetType string `json:"asset_type,omitempty" xml:"asset_type,omitempty"`
	// 承包商编码
	ContractorCode string `json:"contractor_code,omitempty" xml:"contractor_code,omitempty"`
	// 业务方编码
	EmployerCode string `json:"employer_code,omitempty" xml:"employer_code,omitempty"`
	// 用户在业务方平台的userid
	IdentificationInBelongingEmployer string `json:"identification_in_belonging_employer,omitempty" xml:"identification_in_belonging_employer,omitempty"`
	// 税优模式
	TaxOptimizationMode string `json:"tax_optimization_mode,omitempty" xml:"tax_optimization_mode,omitempty"`
}

var poolTaxOptimizationEmployeeAssetUpdateDto = sync.Pool{
	New: func() any {
		return new(TaxOptimizationEmployeeAssetUpdateDto)
	},
}

// GetTaxOptimizationEmployeeAssetUpdateDto() 从对象池中获取TaxOptimizationEmployeeAssetUpdateDto
func GetTaxOptimizationEmployeeAssetUpdateDto() *TaxOptimizationEmployeeAssetUpdateDto {
	return poolTaxOptimizationEmployeeAssetUpdateDto.Get().(*TaxOptimizationEmployeeAssetUpdateDto)
}

// ReleaseTaxOptimizationEmployeeAssetUpdateDto 释放TaxOptimizationEmployeeAssetUpdateDto
func ReleaseTaxOptimizationEmployeeAssetUpdateDto(v *TaxOptimizationEmployeeAssetUpdateDto) {
	v.AssetSymbol = ""
	v.AssetType = ""
	v.ContractorCode = ""
	v.EmployerCode = ""
	v.IdentificationInBelongingEmployer = ""
	v.TaxOptimizationMode = ""
	poolTaxOptimizationEmployeeAssetUpdateDto.Put(v)
}
