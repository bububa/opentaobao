package jst

// MenuInfoReportRequest 
type MenuInfoReportRequest struct {
    // 中间菜单的URL
    MidMenuUrl   string `json:"mid_menu_url,omitempty" xml:"mid_menu_url,omitempty"`
    // 右菜单按钮文案 会员中心、入会有礼、新品推荐、最新活动、店铺微淘
    RightMenuTopic   string `json:"right_menu_topic,omitempty" xml:"right_menu_topic,omitempty"`
    // 左菜单按钮的URL,菜单url限制域名为taobao.com、tmall.com，tb.cn, 且只有180天内有效、用长链。聚石塔侧直接传tbopen链接
    LeftMenuUrl   string `json:"left_menu_url,omitempty" xml:"left_menu_url,omitempty"`
    // 右菜单的URL
    RightMenuUrl   string `json:"right_menu_url,omitempty" xml:"right_menu_url,omitempty"`
    // 中间菜单按钮文案 热销商品、精选商品
    MidMenuTopic   string `json:"mid_menu_topic,omitempty" xml:"mid_menu_topic,omitempty"`
    // ADD：新增，MODIFY：修改
    Operation   string `json:"operation,omitempty" xml:"operation,omitempty"`
    // 左菜单按钮文案 店铺首页
    LeftMenuTopic   string `json:"left_menu_topic,omitempty" xml:"left_menu_topic,omitempty"`
}
