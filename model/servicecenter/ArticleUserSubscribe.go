package servicecenter

// ArticleUserSubscribe 
type ArticleUserSubscribe struct {

    // 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
    ItemCode   string `json:"item_code,omitempty"`

    // 订购关系到期时间
    Deadline   string `json:"deadline,omitempty"`

}
