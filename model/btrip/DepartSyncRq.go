package btrip

// DepartSyncRq 
type DepartSyncRq struct {

    // 部门状态
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 上一版本第三方部门父部门ID（和third_depart_pid字段只传一个即可）
    
    DepartPid   int64 `json:"depart_pid,omitempty" xml:"depart_pid,omitempty"`
    

    // 部门名称
    
    DepartName   string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
    

    // 上一版本第三方部门ID（和third_depart_id只传一个即可）
    
    DepartId   int64 `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
    

    // 新版本第三方部门ID（和depart_id只传一个即可）
    
    ThirdDepartId   string `json:"third_depart_id,omitempty" xml:"third_depart_id,omitempty"`
    

    // 新版本第三方部门父ID（和depart_pid只传一个即可）
    
    ThirdDepartPid   string `json:"third_depart_pid,omitempty" xml:"third_depart_pid,omitempty"`
    

}
