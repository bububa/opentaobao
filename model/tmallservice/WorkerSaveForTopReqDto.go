package tmallservice

// WorkerSaveForTopReqDto 
/* model for simplify = false
type WorkerSaveForTopReqDto struct {

    // 身份证号码
    
    IdNumber   string `json:"id_number,omitempty"`
    

    // 用户地址
    
    Address   string `json:"address,omitempty"`
    

    // 用户地址编码
    
    AddressId   int64 `json:"address_id,omitempty"`
    

    // 真实姓名
    
    RealName   string `json:"real_name,omitempty"`
    

    // 工人技能参数
    
    WorkerServiceAbility  *struct {
        WorkerServiceAbility  *WorkerServiceAbility `json:"worker_service_ability,omitempty"`
    } `json:"worker_service_ability,omitempty"`
    

    // 加入的网点参数
    
    JoinedStore  *struct {
        JoinedStore  *JoinedStore `json:"joined_store,omitempty"`
    } `json:"joined_store,omitempty"`
    

    // 手机号
    
    Phone   string `json:"phone,omitempty"`
    

}
*/

// WorkerSaveForTopReqDto 
type WorkerSaveForTopReqDto struct {

    // 身份证号码
    IdNumber   string `json:"id_number,omitempty"`

    // 用户地址
    Address   string `json:"address,omitempty"`

    // 用户地址编码
    AddressId   int64 `json:"address_id,omitempty"`

    // 真实姓名
    RealName   string `json:"real_name,omitempty"`

    // 工人技能参数
    WorkerServiceAbility   *WorkerServiceAbility `json:"worker_service_ability,omitempty"`

    // 加入的网点参数
    JoinedStore   *JoinedStore `json:"joined_store,omitempty"`

    // 手机号
    Phone   string `json:"phone,omitempty"`

}
