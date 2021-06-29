package tvpay

// GetAppInfoResultDO 
type GetAppInfoResultDO struct {
    // 包名
    PackageName   string `json:"package_name,omitempty" xml:"package_name,omitempty"`
    // 商户id
    PartnerId   int64 `json:"partner_id,omitempty" xml:"partner_id,omitempty"`
    // 应用配置
    AppConfig   *SdkAppConfigDO `json:"app_config,omitempty" xml:"app_config,omitempty"`
    // 全局配置
    GlobalConfig   *SdkGlobalConfigDO `json:"global_config,omitempty" xml:"global_config,omitempty"`
}
