package security

// UrlScanResultItem 
type UrlScanResultItem struct {
    // 风险类型的描述文字
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    // 用户传入的flag
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`
    // 被仿冒的官方网址
    Official   string `json:"official,omitempty" xml:"official,omitempty"`
    // 扫描结果码 SAFE-安全 NOT_SAFE-不安全 UN_KNOWN-未知
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // url风险类型 Phishing-钓鱼网站 Malware-恶意程序 Porn-色情网站 Gambling-赌博网站 Illegal-违法网站 Other-其他
    RiskType   string `json:"risk_type,omitempty" xml:"risk_type,omitempty"`
    // 黑名单来源
    Source   string `json:"source,omitempty" xml:"source,omitempty"`
    // 被钓鱼网站仿冒的对象
    Target   string `json:"target,omitempty" xml:"target,omitempty"`
}
