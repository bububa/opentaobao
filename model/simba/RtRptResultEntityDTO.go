package simba

// RtRptResultEntityDTO 
type RtRptResultEntityDTO struct {

    // 日期
    Thedate   string `json:"thedate,omitempty"`

    // 直接成交金额
    Directtransaction   string `json:"directtransaction,omitempty"`

    // 间接成交金额
    Indirecttransaction   string `json:"indirecttransaction,omitempty"`

    // 直接成交笔数
    Directtransactionshipping   string `json:"directtransactionshipping,omitempty"`

    // 间接成交笔数
    Indirecttransactionshipping   string `json:"indirecttransactionshipping,omitempty"`

    // 收藏宝贝数
    Favitemtotal   string `json:"favitemtotal,omitempty"`

    // 收藏店铺数
    Favshoptotal   string `json:"favshoptotal,omitempty"`

    // 投入产出比
    Roi   string `json:"roi,omitempty"`

    // 点击转化率
    Coverage   string `json:"coverage,omitempty"`

    // 直接购物车数
    Directcarttotal   string `json:"directcarttotal,omitempty"`

    // 间接购物车数
    Indirectcarttotal   string `json:"indirectcarttotal,omitempty"`

    // 购物车总数
    Carttotal   string `json:"carttotal,omitempty"`

    // impression
    Impression   string `json:"impression,omitempty"`

    // cost
    Cost   string `json:"cost,omitempty"`

    // click
    Click   string `json:"click,omitempty"`

    // ctr
    Ctr   string `json:"ctr,omitempty"`

    // cpc
    Cpc   string `json:"cpc,omitempty"`

    // cpm
    Cpm   string `json:"cpm,omitempty"`

    // 总成交笔数
    Transactionshippingtotal   string `json:"transactionshippingtotal,omitempty"`

    // 总成交金额
    Transactiontotal   string `json:"transactiontotal,omitempty"`

    // 总收藏数
    Favtotal   string `json:"favtotal,omitempty"`

    // 计划id
    Campaignid   string `json:"campaignid,omitempty"`

    // 推广组id
    Adgroupid   string `json:"adgroupid,omitempty"`

    // 流量来源( PC站内:1， PC站外:2 , 移动站内:4，移动站外:5)
    Source   string `json:"source,omitempty"`

    // 报表类型（搜索：0,类目出价：1, 单品定向：2, 店铺定向:3）
    SearchType   string `json:"search_type,omitempty"`

    // 人群id
    Crowdid   string `json:"crowdid,omitempty"`

    // 人群名称
    Crowdtitle   string `json:"crowdtitle,omitempty"`

    // 关键词id
    Bidwordid   string `json:"bidwordid,omitempty"`

    // 创意id
    Creativeid   string `json:"creativeid,omitempty"`

}
