package alsc

// CustomerOpenInfo 
/* model for simplify = false
type CustomerOpenInfo struct {

    // 地址
    
    Address   string `json:"address,omitempty"`
    

    // 客单价，单位：分
    
    AvgConsume   int64 `json:"avg_consume,omitempty"`
    

    // 生日
    
    Birthday   string `json:"birthday,omitempty"`
    

    // 渠道
    
    Channel   int64 `json:"channel,omitempty"`
    

    // 累计消费总金额，单位：分
    
    ConsumeAmount   int64 `json:"consume_amount,omitempty"`
    

    // 累计消费次数
    
    ConsumeNum   int64 `json:"consume_num,omitempty"`
    

    // 创建人
    
    CreateBy   string `json:"create_by,omitempty"`
    

    // 顾客ID
    
    CustomerId   string `json:"customer_id,omitempty"`
    

    // 顾客加入门店id
    
    CustomerStoreId   string `json:"customer_store_id,omitempty"`
    

    // 成为顾客时间
    
    CustomerTime   string `json:"customer_time,omitempty"`
    

    // 顾客类型，顾客：0，会员：1
    
    CustomerType   int64 `json:"customer_type,omitempty"`
    

    // 是否已删除
    
    Deleted   bool `json:"deleted,omitempty"`
    

    // 电子邮件
    
    Email   string `json:"email,omitempty"`
    

    // 关注微信时间(微信公众号)
    
    FollowWxTime   string `json:"follow_wx_time,omitempty"`
    

    // 性别  0女 1男
    
    Gender   int64 `json:"gender,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 发票抬头
    
    Invoice   string `json:"invoice,omitempty"`
    

    // 最后消费时间
    
    LastConsumeTime   string `json:"last_consume_time,omitempty"`
    

    // 等级描述
    
    LevelDesc   string `json:"level_desc,omitempty"`
    

    // 等级ID
    
    LevelId   string `json:"level_id,omitempty"`
    

    // 等级序号
    
    LevelNo   int64 `json:"level_no,omitempty"`
    

    // 成为会员店铺ID
    
    MemberStoreId   string `json:"member_store_id,omitempty"`
    

    // 成为会员时间
    
    MemberTime   string `json:"member_time,omitempty"`
    

    // 手机号
    
    Mobile   string `json:"mobile,omitempty"`
    

    // 姓名
    
    Name   string `json:"name,omitempty"`
    

    // 固定电话
    
    Phone   string `json:"phone,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

    // 会员状态，可用1、停用0(停用：所有会员的权益不可用)
    
    State   int64 `json:"state,omitempty"`
    

    // 标签列表
    
    Tags  struct {
        String  []string `json:"string,omitempty"`
    } `json:"tags,omitempty"`
    

    // 2019237428364
    
    UpdateBy   string `json:"update_by,omitempty"`
    

    // 成长值
    
    Growth   int64 `json:"growth,omitempty"`
    

    // 是否设置了支付密码
    
    HasPassword   bool `json:"has_password,omitempty"`
    

    // 外部信息
    
    CustomerOutInfoList  struct {
        CustomerOutInfo  []CustomerOutInfo `json:"customer_out_info,omitempty"`
    } `json:"customer_out_info_list,omitempty"`
    

}
*/

// CustomerOpenInfo 
type CustomerOpenInfo struct {

    // 地址
    Address   string `json:"address,omitempty"`

    // 客单价，单位：分
    AvgConsume   int64 `json:"avg_consume,omitempty"`

    // 生日
    Birthday   string `json:"birthday,omitempty"`

    // 渠道
    Channel   int64 `json:"channel,omitempty"`

    // 累计消费总金额，单位：分
    ConsumeAmount   int64 `json:"consume_amount,omitempty"`

    // 累计消费次数
    ConsumeNum   int64 `json:"consume_num,omitempty"`

    // 创建人
    CreateBy   string `json:"create_by,omitempty"`

    // 顾客ID
    CustomerId   string `json:"customer_id,omitempty"`

    // 顾客加入门店id
    CustomerStoreId   string `json:"customer_store_id,omitempty"`

    // 成为顾客时间
    CustomerTime   string `json:"customer_time,omitempty"`

    // 顾客类型，顾客：0，会员：1
    CustomerType   int64 `json:"customer_type,omitempty"`

    // 是否已删除
    Deleted   bool `json:"deleted,omitempty"`

    // 电子邮件
    Email   string `json:"email,omitempty"`

    // 关注微信时间(微信公众号)
    FollowWxTime   string `json:"follow_wx_time,omitempty"`

    // 性别  0女 1男
    Gender   int64 `json:"gender,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 发票抬头
    Invoice   string `json:"invoice,omitempty"`

    // 最后消费时间
    LastConsumeTime   string `json:"last_consume_time,omitempty"`

    // 等级描述
    LevelDesc   string `json:"level_desc,omitempty"`

    // 等级ID
    LevelId   string `json:"level_id,omitempty"`

    // 等级序号
    LevelNo   int64 `json:"level_no,omitempty"`

    // 成为会员店铺ID
    MemberStoreId   string `json:"member_store_id,omitempty"`

    // 成为会员时间
    MemberTime   string `json:"member_time,omitempty"`

    // 手机号
    Mobile   string `json:"mobile,omitempty"`

    // 姓名
    Name   string `json:"name,omitempty"`

    // 固定电话
    Phone   string `json:"phone,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 会员状态，可用1、停用0(停用：所有会员的权益不可用)
    State   int64 `json:"state,omitempty"`

    // 标签列表
    Tags   []string `json:"tags,omitempty"`

    // 2019237428364
    UpdateBy   string `json:"update_by,omitempty"`

    // 成长值
    Growth   int64 `json:"growth,omitempty"`

    // 是否设置了支付密码
    HasPassword   bool `json:"has_password,omitempty"`

    // 外部信息
    CustomerOutInfoList   []CustomerOutInfo `json:"customer_out_info_list,omitempty"`

}
