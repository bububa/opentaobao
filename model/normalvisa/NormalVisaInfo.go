package normalvisa

// NormalVisaInfo 
type NormalVisaInfo struct {
    // 国家id，国家编码详见：https://open.alitrip.com/docs/doc.htm?spm=a21tt.7629140.0.0.retXmq&treeId=79&articleId=104840&docType=1
    CountryId   string `json:"country_id,omitempty" xml:"country_id,omitempty"`
    // 1:贴纸签  2:电子签  3:面试
    NormalVisaType   int64 `json:"normal_visa_type,omitempty" xml:"normal_visa_type,omitempty"`
    // 无办理人信息：1001 办理人已填写：1002 已收到资料: 1003 已审核完成: 1004 已送签:1005 结果已返回: 1006 已预约面试: 1007 处理中:1008 已中止办理：1010
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 支付时间
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    // 数量
    Amount   int64 `json:"amount,omitempty" xml:"amount,omitempty"`
    // 单价
    AuctionPrice   string `json:"auction_price,omitempty" xml:"auction_price,omitempty"`
    // 标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 卖家昵称
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    // 是否达到中止状态
    EndStatus   bool `json:"end_status,omitempty" xml:"end_status,omitempty"`
    // 买家昵称
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
    // 订单号
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    // 无办理人信息：1001 办理人已填写：1002 已收到资料: 1003 已审核完成: 1004 已送签:1005 结果已返回: 1006 已预约面试: 1007 处理中:1008 已中止办理：1010
    StatusDesc   string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
    // 是否需要商家代填
    NeedFillContact   bool `json:"need_fill_contact,omitempty" xml:"need_fill_contact,omitempty"`
}
