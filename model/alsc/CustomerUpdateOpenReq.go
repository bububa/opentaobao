package alsc

// CustomerUpdateOpenReq 
/* model for simplify = false
type CustomerUpdateOpenReq struct {

    // 顾客ID
    
    CustomerId   string `json:"customer_id,omitempty"`
    

    // 品牌ID 外部品牌id  2选1
    
    BrandId   string `json:"brand_id,omitempty"`
    

    // 地址
    
    Address   string `json:"address,omitempty"`
    

    // 生日
    
    Birthday   string `json:"birthday,omitempty"`
    

    // 邮箱
    
    Email   string `json:"email,omitempty"`
    

    // 扩展信息
    
    ExtInfo   string `json:"ext_info,omitempty"`
    

    // 性别
    
    Gender   int64 `json:"gender,omitempty"`
    

    // 发票抬头
    
    Invoice   string `json:"invoice,omitempty"`
    

    // 等级ID
    
    LevelId   string `json:"level_id,omitempty"`
    

    // 移动电话
    
    Mobile   string `json:"mobile,omitempty"`
    

    // 姓名
    
    Name   string `json:"name,omitempty"`
    

    // 操作人ID
    
    OperatorId   string `json:"operator_id,omitempty"`
    

    // 操作人姓名
    
    OperatorName   string `json:"operator_name,omitempty"`
    

    // 固定电话
    
    Phone   string `json:"phone,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

    // 请求ID，幂等处理
    
    RequestId   string `json:"request_id,omitempty"`
    

    // 人群标签
    
    TagIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"tag_ids,omitempty"`
    

    // 外部品牌id
    
    OutBrandId   string `json:"out_brand_id,omitempty"`
    

    // 0顾客，1会员
    
    CustomerType   int64 `json:"customer_type,omitempty"`
    

}
*/

// CustomerUpdateOpenReq 
type CustomerUpdateOpenReq struct {

    // 顾客ID
    CustomerId   string `json:"customer_id,omitempty"`

    // 品牌ID 外部品牌id  2选1
    BrandId   string `json:"brand_id,omitempty"`

    // 地址
    Address   string `json:"address,omitempty"`

    // 生日
    Birthday   string `json:"birthday,omitempty"`

    // 邮箱
    Email   string `json:"email,omitempty"`

    // 扩展信息
    ExtInfo   string `json:"ext_info,omitempty"`

    // 性别
    Gender   int64 `json:"gender,omitempty"`

    // 发票抬头
    Invoice   string `json:"invoice,omitempty"`

    // 等级ID
    LevelId   string `json:"level_id,omitempty"`

    // 移动电话
    Mobile   string `json:"mobile,omitempty"`

    // 姓名
    Name   string `json:"name,omitempty"`

    // 操作人ID
    OperatorId   string `json:"operator_id,omitempty"`

    // 操作人姓名
    OperatorName   string `json:"operator_name,omitempty"`

    // 固定电话
    Phone   string `json:"phone,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 请求ID，幂等处理
    RequestId   string `json:"request_id,omitempty"`

    // 人群标签
    TagIds   []string `json:"tag_ids,omitempty"`

    // 外部品牌id
    OutBrandId   string `json:"out_brand_id,omitempty"`

    // 0顾客，1会员
    CustomerType   int64 `json:"customer_type,omitempty"`

}
