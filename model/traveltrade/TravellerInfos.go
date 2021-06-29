package traveltrade

// TravellerInfos 
type TravellerInfos struct {

    // 证件签发日期
    
    IssuedDate   string `json:"issued_date,omitempty" xml:"issued_date,omitempty"`
    

    // 拼音姓
    
    LastName   string `json:"last_name,omitempty" xml:"last_name,omitempty"`
    

    // 1： 出行管理 2： 二次预约 4： 二次确认 8:   电子合同 16：电子凭证 32：自定义出行人 64：门票
    
    BizType   string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    

    // 预约ID
    
    BookInfoId   int64 `json:"book_info_id,omitempty" xml:"book_info_id,omitempty"`
    

    // 子订单ID
    
    SubTcOrderId   string `json:"sub_tc_order_id,omitempty" xml:"sub_tc_order_id,omitempty"`
    

    // 子订单ID
    
    TcOrderId   string `json:"tc_order_id,omitempty" xml:"tc_order_id,omitempty"`
    

    // 证件类型
    
    CredentialsType   int64 `json:"credentials_type,omitempty" xml:"credentials_type,omitempty"`
    

    // 出行人姓名
    
    TravellerName   string `json:"traveller_name,omitempty" xml:"traveller_name,omitempty"`
    

    // 买家备注
    
    BuyerRemark   string `json:"buyer_remark,omitempty" xml:"buyer_remark,omitempty"`
    

    // 修改时间
    
    ModifyTime   string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
    

    // 类型 出行人（签署人）为0或空，联系人为1
    
    PersonType   string `json:"person_type,omitempty" xml:"person_type,omitempty"`
    

    // sku
    
    Sku   string `json:"sku,omitempty" xml:"sku,omitempty"`
    

    // 证件号码
    
    CredentialsCode   string `json:"credentials_code,omitempty" xml:"credentials_code,omitempty"`
    

    // 根据子订单ID列表查询
    
    SubTcOrderIds   string `json:"sub_tc_order_ids,omitempty" xml:"sub_tc_order_ids,omitempty"`
    

    // 姓名的拼音，格式为：xing/ming,中间以“/”分隔
    
    TravellerNamePinyin   string `json:"traveller_name_pinyin,omitempty" xml:"traveller_name_pinyin,omitempty"`
    

    // 用户是通过哪个终端填写
    
    TerminalType   string `json:"terminal_type,omitempty" xml:"terminal_type,omitempty"`
    

    // 拼音名
    
    FirstName   string `json:"first_name,omitempty" xml:"first_name,omitempty"`
    

    // 商品ID
    
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 国籍
    
    Nationality   string `json:"nationality,omitempty" xml:"nationality,omitempty"`
    

    // 是否仅更新bookId字段
    
    OnlyUpateBookId   bool `json:"only_upate_book_id,omitempty" xml:"only_upate_book_id,omitempty"`
    

    // 创建时间
    
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    

    // 证件签发国
    
    IssuedCountry   string `json:"issued_country,omitempty" xml:"issued_country,omitempty"`
    

    // 电话
    
    Telphone   string `json:"telphone,omitempty" xml:"telphone,omitempty"`
    

    // 自定义出行人的时候的json格式
    
    TravellerJson   string `json:"traveller_json,omitempty" xml:"traveller_json,omitempty"`
    

}
