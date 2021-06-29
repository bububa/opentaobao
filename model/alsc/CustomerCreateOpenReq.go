package alsc

// CustomerCreateOpenReq 
type CustomerCreateOpenReq struct {
    // 外部ID
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    // 外部ID类型
    OuterType   string `json:"outer_type,omitempty" xml:"outer_type,omitempty"`
    // 会员地址
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    // 生日
    Birthday   string `json:"birthday,omitempty" xml:"birthday,omitempty"`
    // 品牌ID  外部品牌id
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    // 渠道
    Channel   string `json:"channel,omitempty" xml:"channel,omitempty"`
    // 顾客类型，1:会员，0:顾客
    CustomerType   int64 `json:"customer_type,omitempty" xml:"customer_type,omitempty"`
    // 邮箱
    Email   string `json:"email,omitempty" xml:"email,omitempty"`
    // 性别  0女 1男
    Gender   string `json:"gender,omitempty" xml:"gender,omitempty"`
    // 发票
    Invoice   string `json:"invoice,omitempty" xml:"invoice,omitempty"`
    // 等级ID
    LevelId   string `json:"level_id,omitempty" xml:"level_id,omitempty"`
    // 手机号
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    // 姓名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 操作人ID
    OperatorId   string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
    // 固定电话
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 请求ID，用于保障幂等性
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 店铺ID
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    // 标签列表
    TagIdList   []string `json:"tag_id_list,omitempty" xml:"tag_id_list>string,omitempty"`
    // 外部品牌id
    OutBrandId   string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
    // 外部门店id
    OutShopId   string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
}
