package simba

// ResultMap 
type ResultMap struct {

    // ctr
    Ctr   string `json:"ctr,omitempty"`

    // cpm
    Cpm   string `json:"cpm,omitempty"`

    // 消耗（分）
    Cost   string `json:"cost,omitempty"`

    // 计划id
    Campaignid   string `json:"campaignid,omitempty"`

    // 词包名称（汇总-流量包、汇总-自主买词）
    SubPackageName   string `json:"sub_package_name,omitempty"`

    // 点击量
    Click   string `json:"click,omitempty"`

    // 词包类型（0：汇总-流量包，-1：汇总-自主买词）
    SubPackage   string `json:"sub_package,omitempty"`

    // 报表数据日期
    Thedate   string `json:"thedate,omitempty"`

    // 推广组id
    Adgroupid   string `json:"adgroupid,omitempty"`

    // 展现量
    Impression   string `json:"impression,omitempty"`

    // 点击转化率
    Coverage   string `json:"coverage,omitempty"`

    // roi
    Roi   string `json:"roi,omitempty"`

}
