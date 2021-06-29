package alihealth2

// IsvRiseReserveRequest 
type IsvRiseReserveRequest struct {

    // 证件类型 0身份证 1护照
    
    IdType   int64 `json:"id_type,omitempty" xml:"id_type,omitempty"`
    

    // 支付宝订单ID
    
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 性别 0男1女
    
    Sex   int64 `json:"sex,omitempty" xml:"sex,omitempty"`
    

    // ISV门店名称
    
    SpStoreName   string `json:"sp_store_name,omitempty" xml:"sp_store_name,omitempty"`
    

    // 证件号码
    
    IdNumber   string `json:"id_number,omitempty" xml:"id_number,omitempty"`
    

    // ISV项目名称，多个项目英文逗号分割
    
    SpItemName   string `json:"sp_item_name,omitempty" xml:"sp_item_name,omitempty"`
    

    // ISV项目ID,多个项目英文逗号分割
    
    SpItemId   string `json:"sp_item_id,omitempty" xml:"sp_item_id,omitempty"`
    

    // 用户号码
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    

    // 用户姓名
    
    UserNick   string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
    

    // ISV门店ID
    
    SpStoreId   string `json:"sp_store_id,omitempty" xml:"sp_store_id,omitempty"`
    

    // 预约时间
    
    ReserveTime   string `json:"reserve_time,omitempty" xml:"reserve_time,omitempty"`
    

    // 是否结婚 0未婚1结婚
    
    IsMarried   int64 `json:"is_married,omitempty" xml:"is_married,omitempty"`
    

    // 支付宝标品ID
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // ISV复诊/升单ID
    
    SpReserveId   string `json:"sp_reserve_id,omitempty" xml:"sp_reserve_id,omitempty"`
    

    // 支付宝初诊ID
    
    FirstReserveId   int64 `json:"first_reserve_id,omitempty" xml:"first_reserve_id,omitempty"`
    

    // userId
    
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

    // 科室名称
    
    DepartmentName   string `json:"department_name,omitempty" xml:"department_name,omitempty"`
    

    // 0 新增 1 修改
    
    OpType   int64 `json:"op_type,omitempty" xml:"op_type,omitempty"`
    

    // 医生名称
    
    DoctorName   string `json:"doctor_name,omitempty" xml:"doctor_name,omitempty"`
    

    // 修改状态 2 待就诊 4 已就诊
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

}
