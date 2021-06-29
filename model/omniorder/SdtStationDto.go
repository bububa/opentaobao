package omniorder

// SdtStationDto 
type SdtStationDto struct {

    // 站点操作时间
    
    ActionTime   string `json:"action_time,omitempty" xml:"action_time,omitempty"`
    

    // 快递公司cpcode
    
    CpCode   string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
    

    // 快递公司名称
    
    CpName   string `json:"cp_name,omitempty" xml:"cp_name,omitempty"`
    

    // 站点code
    
    StationCode   string `json:"station_code,omitempty" xml:"station_code,omitempty"`
    

    // 站点联系方式
    
    StationContact   string `json:"station_contact,omitempty" xml:"station_contact,omitempty"`
    

    // 站点负责人
    
    StationMaster   string `json:"station_master,omitempty" xml:"station_master,omitempty"`
    

    // 站点名
    
    StationName   string `json:"station_name,omitempty" xml:"station_name,omitempty"`
    

    // 站点类别（推荐站点、派送站点、揽收站点）
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

}
