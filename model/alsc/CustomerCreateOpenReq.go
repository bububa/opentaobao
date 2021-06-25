package alsc

// CustomerCreateOpenReq 
type CustomerCreateOpenReq struct {

    // 外部ID
    OuterId   string `json:"outer_id,omitempty"`

    // 外部ID类型
    OuterType   string `json:"outer_type,omitempty"`

    // 会员地址
    Address   string `json:"address,omitempty"`

    // 生日
    Birthday   string `json:"birthday,omitempty"`

    // 品牌ID  外部品牌id
    BrandId   string `json:"brand_id,omitempty"`

    // 渠道
    Channel   string `json:"channel,omitempty"`

    // 顾客类型，1:会员，0:顾客
    CustomerType   int64 `json:"customer_type,omitempty"`

    // 邮箱
    Email   string `json:"email,omitempty"`

    // 性别  0女 1男
    Gender   string `json:"gender,omitempty"`

    // 发票
    Invoice   string `json:"invoice,omitempty"`

    // 等级ID
    LevelId   string `json:"level_id,omitempty"`

    // 手机号
    Mobile   string `json:"mobile,omitempty"`

    // 姓名
    Name   string `json:"name,omitempty"`

    // 操作人ID
    OperatorId   string `json:"operator_id,omitempty"`

    // 固定电话
    Phone   string `json:"phone,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 请求ID，用于保障幂等性
    RequestId   string `json:"request_id,omitempty"`

    // 店铺ID
    ShopId   string `json:"shop_id,omitempty"`

    // 标签列表
    TagIdList   []String `json:"tag_id_list,omitempty"`

    // 外部品牌id
    OutBrandId   string `json:"out_brand_id,omitempty"`

    // 外部门店id
    OutShopId   string `json:"out_shop_id,omitempty"`

}
