package ticket

// TopTicketItemFullinfoResult 
type TopTicketItemFullinfoResult struct {

    // 阿里旅行提供的景点编码
    AliScenicId   int64 `json:"ali_scenic_id,omitempty"`

    // 联票的场景，则该收费项目可能关联多个景点，以英文逗号分隔
    AliScenicIds   string `json:"ali_scenic_ids,omitempty"`

    // 商户系统中景点编码
    OutScenicId   string `json:"out_scenic_id,omitempty"`

    // 阿里旅行提供的收费项目编码
    AliProductId   int64 `json:"ali_product_id,omitempty"`

    // 阿里旅行收费项目名称
    AliProductName   string `json:"ali_product_name,omitempty"`

    // 商户自定义收费项目编码
    OutProductId   string `json:"out_product_id,omitempty"`

    // 商户收费项目名称
    OutProductName   string `json:"out_product_name,omitempty"`

    // 门票商品 库存类型。1、每日库存， 2、区间总库存
    InventoryType   int64 `json:"inventory_type,omitempty"`

    // 是否需要买家指定入园日期。1、需要，2-不需要
    NeedEnterDate   int64 `json:"need_enter_date,omitempty"`

    // 门票有效期：指定入园日期后 多少天内有效。当为数字时，表示多少天内有效；当为日期时，表示到某日期有效，日期格式：yyyy-MM-dd
    ExpireDate   string `json:"expire_date,omitempty"`

    // 门票 预定时间限制。1、表示无限制 购买后可立即入园，2、有限制，此时预定时间限制规则必填。
    ReserveLimitType   int64 `json:"reserve_limit_type,omitempty"`

    // 门票 预定时间限制规则。格式：1_18_00_3，含义：必须提前1天拍下，且在18点00分前支付成功，订单才生效。当为提前0天时（即当日票），最后一个数字才有用，指当日票需要在预定3小时后入园。
    ReserveLimitRule   string `json:"reserve_limit_rule,omitempty"`

    // 门票商品发码方式
    CodeSendingInfo   *CodeSendingInfo `json:"code_sending_info,omitempty"`

    // 商品ID
    ItemId   int64 `json:"item_id,omitempty"`

    // 商品标题
    Title   string `json:"title,omitempty"`

    // 商品主图
    PicUrls   []String `json:"pic_urls,omitempty"`

    // 商品详情描述
    Desc   string `json:"desc,omitempty"`

    // 手机描述
    WapDesc   string `json:"wap_desc,omitempty"`

    // 商品状态 0-下架，1-上架
    ItemStatus   int64 `json:"item_status,omitempty"`

    // 门票商品下 各个票种的sku信息
    TicketTypes   []TicketSimpleSkuParam `json:"ticket_types,omitempty"`

}
