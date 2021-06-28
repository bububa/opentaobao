package security

// FakeAppSummary 
/* model for simplify = false
type FakeAppSummary struct {

    // 仿冒应用总数(任务完成时才返回)
    
    FakeAppCount   int64 `json:"fake_app_count,omitempty"`
    

    // 子任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
    
    Status   int64 `json:"status,omitempty"`
    

    // 仿冒应用感染用户总数(任务完成时才返回)
    
    TotalInfectedUsers   int64 `json:"total_infected_users,omitempty"`
    

}
*/

// FakeAppSummary 
type FakeAppSummary struct {

    // 仿冒应用总数(任务完成时才返回)
    FakeAppCount   int64 `json:"fake_app_count,omitempty"`

    // 子任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
    Status   int64 `json:"status,omitempty"`

    // 仿冒应用感染用户总数(任务完成时才返回)
    TotalInfectedUsers   int64 `json:"total_infected_users,omitempty"`

}
