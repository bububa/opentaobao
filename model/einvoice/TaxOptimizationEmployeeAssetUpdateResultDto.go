package einvoice

// TaxOptimizationEmployeeAssetUpdateResultDto 
type TaxOptimizationEmployeeAssetUpdateResultDto struct {

    // 承包商编码
    
    ContractorCode   string `json:"contractor_code,omitempty" xml:"contractor_code,omitempty"`
    

    // 需要更新的资产账号
    
    CurrentBindedAssetSymbol   string `json:"current_binded_asset_symbol,omitempty" xml:"current_binded_asset_symbol,omitempty"`
    

    // 需要更新的资产类型
    
    CurrentBindedAssetType   string `json:"current_binded_asset_type,omitempty" xml:"current_binded_asset_type,omitempty"`
    

    // 业务方编码
    
    EmployerCode   string `json:"employer_code,omitempty" xml:"employer_code,omitempty"`
    

    // 用户在业务方平台的userid
    
    IdentificationInBelongingEmployer   string `json:"identification_in_belonging_employer,omitempty" xml:"identification_in_belonging_employer,omitempty"`
    

    // 当前资产账号
    
    PreviousBindedAssetSymbol   string `json:"previous_binded_asset_symbol,omitempty" xml:"previous_binded_asset_symbol,omitempty"`
    

    // 当前资产类型
    
    PreviousBindedAssetType   string `json:"previous_binded_asset_type,omitempty" xml:"previous_binded_asset_type,omitempty"`
    

    // 税优模式
    
    TaxOptimizationMode   string `json:"tax_optimization_mode,omitempty" xml:"tax_optimization_mode,omitempty"`
    

}
