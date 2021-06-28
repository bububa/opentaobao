package eticket

// PackageResult 
/* model for simplify = false
type PackageResult struct {

    // 操作结果码
    
    Code   int64 `json:"code,omitempty"`
    

    // 操作是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 操作结果信息
    
    Info   string `json:"info,omitempty"`
    

    // 包基本信息
    
    PackageBase  *struct {
        PackageBase  *PackageBase `json:"package_base,omitempty"`
    } `json:"package_base,omitempty"`
    

    // 包基本信息列表
    
    PackageBaseList  struct {
        PackageBase  []PackageBase `json:"package_base,omitempty"`
    } `json:"package_base_list,omitempty"`
    

}
*/

// PackageResult 
type PackageResult struct {

    // 操作结果码
    Code   int64 `json:"code,omitempty"`

    // 操作是否成功
    Success   bool `json:"success,omitempty"`

    // 操作结果信息
    Info   string `json:"info,omitempty"`

    // 包基本信息
    PackageBase   *PackageBase `json:"package_base,omitempty"`

    // 包基本信息列表
    PackageBaseList   []PackageBase `json:"package_base_list,omitempty"`

}
