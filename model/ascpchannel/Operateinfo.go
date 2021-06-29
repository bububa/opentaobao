package ascpchannel

// Operateinfo 
type Operateinfo struct {

    // 操作时间Date类型
    
    OperateTime   string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
    

    // 操作人
    
    OperateName   string `json:"operate_name,omitempty" xml:"operate_name,omitempty"`
    

}
