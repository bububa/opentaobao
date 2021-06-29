package idleparttime

// PartTimeJobSyncLogList 
type PartTimeJobSyncLogList struct {

    // 日志节点
    
    PartTimeJobTransNodes   []PartTimeJobTransNodes `json:"part_time_job_trans_nodes,omitempty" xml:"part_time_job_trans_nodes,omitempty"`
    

    // 岗位id
    
    JobId   int64 `json:"job_id,omitempty" xml:"job_id,omitempty"`
    

    // 当前状态描述
    
    CurrentStatus   string `json:"current_status,omitempty" xml:"current_status,omitempty"`
    

    // 岗位创建时间
    
    CreateTime   int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
    

}
