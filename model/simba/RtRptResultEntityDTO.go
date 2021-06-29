package simba

// RtRptResultEntityDTO 
type RtRptResultEntityDTO struct {
    // 日期
    Thedate   string `json:"thedate,omitempty" xml:"thedate,omitempty"`
    // 直接成交金额
    Directtransaction   string `json:"directtransaction,omitempty" xml:"directtransaction,omitempty"`
    // 间接成交金额
    Indirecttransaction   string `json:"indirecttransaction,omitempty" xml:"indirecttransaction,omitempty"`
    // 直接成交笔数
    Directtransactionshipping   string `json:"directtransactionshipping,omitempty" xml:"directtransactionshipping,omitempty"`
    // 间接成交笔数
    Indirecttransactionshipping   string `json:"indirecttransactionshipping,omitempty" xml:"indirecttransactionshipping,omitempty"`
    // 收藏宝贝数
    Favitemtotal   string `json:"favitemtotal,omitempty" xml:"favitemtotal,omitempty"`
    // 收藏店铺数
    Favshoptotal   string `json:"favshoptotal,omitempty" xml:"favshoptotal,omitempty"`
    // 投入产出比
    Roi   string `json:"roi,omitempty" xml:"roi,omitempty"`
    // 点击转化率
    Coverage   string `json:"coverage,omitempty" xml:"coverage,omitempty"`
    // 直接购物车数
    Directcarttotal   string `json:"directcarttotal,omitempty" xml:"directcarttotal,omitempty"`
    // 间接购物车数
    Indirectcarttotal   string `json:"indirectcarttotal,omitempty" xml:"indirectcarttotal,omitempty"`
    // 购物车总数
    Carttotal   string `json:"carttotal,omitempty" xml:"carttotal,omitempty"`
    // impression
    Impression   string `json:"impression,omitempty" xml:"impression,omitempty"`
    // cost
    Cost   string `json:"cost,omitempty" xml:"cost,omitempty"`
    // click
    Click   string `json:"click,omitempty" xml:"click,omitempty"`
    // ctr
    Ctr   string `json:"ctr,omitempty" xml:"ctr,omitempty"`
    // cpc
    Cpc   string `json:"cpc,omitempty" xml:"cpc,omitempty"`
    // cpm
    Cpm   string `json:"cpm,omitempty" xml:"cpm,omitempty"`
    // 总成交笔数
    Transactionshippingtotal   string `json:"transactionshippingtotal,omitempty" xml:"transactionshippingtotal,omitempty"`
    // 总成交金额
    Transactiontotal   string `json:"transactiontotal,omitempty" xml:"transactiontotal,omitempty"`
    // 总收藏数
    Favtotal   string `json:"favtotal,omitempty" xml:"favtotal,omitempty"`
    // 计划id
    Campaignid   string `json:"campaignid,omitempty" xml:"campaignid,omitempty"`
    // 推广组id
    Adgroupid   string `json:"adgroupid,omitempty" xml:"adgroupid,omitempty"`
    // 流量来源( PC站内:1， PC站外:2 , 移动站内:4，移动站外:5)
    Source   string `json:"source,omitempty" xml:"source,omitempty"`
    // 报表类型（搜索：0,类目出价：1, 单品定向：2, 店铺定向:3）
    SearchType   string `json:"search_type,omitempty" xml:"search_type,omitempty"`
    // 人群id
    Crowdid   string `json:"crowdid,omitempty" xml:"crowdid,omitempty"`
    // 人群名称
    Crowdtitle   string `json:"crowdtitle,omitempty" xml:"crowdtitle,omitempty"`
    // 关键词id
    Bidwordid   string `json:"bidwordid,omitempty" xml:"bidwordid,omitempty"`
    // 创意id
    Creativeid   string `json:"creativeid,omitempty" xml:"creativeid,omitempty"`
}
