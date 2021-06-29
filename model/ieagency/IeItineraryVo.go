package ieagency

// IeItineraryVo 
type IeItineraryVo struct {

    // 快递公司编号
    
    ExpressCompanyName   string `json:"express_company_name,omitempty" xml:"express_company_name,omitempty"`
    

    // 退款金额
    
    RefundMoney   int64 `json:"refund_money,omitempty" xml:"refund_money,omitempty"`
    

    // 行程单状态(INITIAL:初始状态,NOPAY(:未付款,PAYED:已付款,FINISHED:已完成,REFUND:已退款,CLOSED:已关闭（随主订单))
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 商家备注信息
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 收件人
    
    Consignee   string `json:"consignee,omitempty" xml:"consignee,omitempty"`
    

    // 城市名称
    
    CityName   string `json:"city_name,omitempty" xml:"city_name,omitempty"`
    

    // 发送时间
    
    SendTime   string `json:"send_time,omitempty" xml:"send_time,omitempty"`
    

    // 配送方式(Counter:自取,Normal:快递,Regist:挂号信,Common:平邮)
    
    SendType   string `json:"send_type,omitempty" xml:"send_type,omitempty"`
    

    // 省份名称
    
    ProvinceName   string `json:"province_name,omitempty" xml:"province_name,omitempty"`
    

    // 邮编
    
    Postcode   string `json:"postcode,omitempty" xml:"postcode,omitempty"`
    

    // 主键
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

    // 区县名称
    
    DistrictName   string `json:"district_name,omitempty" xml:"district_name,omitempty"`
    

    // 价格
    
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    

    // 收货地址
    
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    

    // 支付状态(INITIAL: 初始状态, CREATED:已创建担保交易, PAYED:已付款, TURNED:转交易成功)
    
    PayStatus   string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
    

    // 快递单号
    
    ExpressNo   string `json:"express_no,omitempty" xml:"express_no,omitempty"`
    

    // 行程单号
    
    ItineraryNo   string `json:"itinerary_no,omitempty" xml:"itinerary_no,omitempty"`
    

    // 电话
    
    Telephone   string `json:"telephone,omitempty" xml:"telephone,omitempty"`
    

    // 是否已经邮寄
    
    IsPost   bool `json:"is_post,omitempty" xml:"is_post,omitempty"`
    

    // 快递公司编号
    
    ExpressCompanyCode   string `json:"express_company_code,omitempty" xml:"express_company_code,omitempty"`
    

    // 移动电话
    
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    

}
