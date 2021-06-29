package servicecenter

// ArticleSub 
type ArticleSub struct {
    // 淘宝会员名
    Nick   string `json:"nick,omitempty" xml:"nick,omitempty"`
    // 应用名称
    ArticleName   string `json:"article_name,omitempty" xml:"article_name,omitempty"`
    // 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
    ArticleCode   string `json:"article_code,omitempty" xml:"article_code,omitempty"`
    // 收费项目名称
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    // 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
    // 订购关系到期时间
    Deadline   string `json:"deadline,omitempty" xml:"deadline,omitempty"`
    // 状态，1=有效 2=过期
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 是否自动续费
    Autosub   bool `json:"autosub,omitempty" xml:"autosub,omitempty"`
    // 是否到期提醒
    ExpireNotice   bool `json:"expire_notice,omitempty" xml:"expire_notice,omitempty"`
}
