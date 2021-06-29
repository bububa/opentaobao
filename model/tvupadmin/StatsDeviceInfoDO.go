package tvupadmin

// StatsDeviceInfoDO 
type StatsDeviceInfoDO struct {

    // id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // factoryName
    
    FactoryName   string `json:"factory_name,omitempty" xml:"factory_name,omitempty"`
    

    // deviceModel
    
    DeviceModel   string `json:"device_model,omitempty" xml:"device_model,omitempty"`
    

    // userActiveTotal
    
    UserActiveTotal   int64 `json:"user_active_total,omitempty" xml:"user_active_total,omitempty"`
    

    // userSignDaily
    
    UserSignDaily   int64 `json:"user_sign_daily,omitempty" xml:"user_sign_daily,omitempty"`
    

    // statsDateStr
    
    StatsDateStr   string `json:"stats_date_str,omitempty" xml:"stats_date_str,omitempty"`
    

    // statsDate
    
    StatsDate   string `json:"stats_date,omitempty" xml:"stats_date,omitempty"`
    

}
