package legalsuit

// CommunicateModel 
type CommunicateModel struct {

    // 沟通背景、沟通目的
    
    CommunicateAim   string `json:"communicate_aim,omitempty" xml:"communicate_aim,omitempty"`
    

    // 联系方式
    
    Contact   string `json:"contact,omitempty" xml:"contact,omitempty"`
    

    // 人员姓名
    
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    

    // 沟通方式
    
    CommunicateType   string `json:"communicate_type,omitempty" xml:"communicate_type,omitempty"`
    

    // 单位名称
    
    UnitName   string `json:"unit_name,omitempty" xml:"unit_name,omitempty"`
    

    // 机构信息
    
    Department   string `json:"department,omitempty" xml:"department,omitempty"`
    

    // 沟通时间
    
    CommunicateTime   string `json:"communicate_time,omitempty" xml:"communicate_time,omitempty"`
    

    // 沟通内容
    
    CommunicateContent   string `json:"communicate_content,omitempty" xml:"communicate_content,omitempty"`
    

    // 下一步动作
    
    Next   string `json:"next,omitempty" xml:"next,omitempty"`
    

}
