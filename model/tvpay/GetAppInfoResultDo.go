package tvpay

// GetAppInfoResultDo 
/* model for simplify = false
type GetAppInfoResultDo struct {

    // 包名
    
    PackageName   string `json:"package_name,omitempty"`
    

    // 商户id
    
    PartnerId   int64 `json:"partner_id,omitempty"`
    

    // 应用配置
    
    AppConfig  *struct {
        SdkAppConfigDo  *SdkAppConfigDo `json:"sdk_app_config_do,omitempty"`
    } `json:"app_config,omitempty"`
    

    // 全局配置
    
    GlobalConfig  *struct {
        SdkGlobalConfigDo  *SdkGlobalConfigDo `json:"sdk_global_config_do,omitempty"`
    } `json:"global_config,omitempty"`
    

}
*/

// GetAppInfoResultDo 
type GetAppInfoResultDo struct {

    // 包名
    PackageName   string `json:"package_name,omitempty"`

    // 商户id
    PartnerId   int64 `json:"partner_id,omitempty"`

    // 应用配置
    AppConfig   *SdkAppConfigDo `json:"app_config,omitempty"`

    // 全局配置
    GlobalConfig   *SdkGlobalConfigDo `json:"global_config,omitempty"`

}
