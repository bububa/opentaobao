package mtopopen

// TopIsvDecorateParam 
/* model for simplify = false
type TopIsvDecorateParam struct {

    // 活动id，调用alibaba.interact.activity.register传入的bizid
    
    BizId   string `json:"biz_id,omitempty"`
    

    // 目前必须填0，代表店铺
    
    BizType   string `json:"biz_type,omitempty"`
    

    // {"action":"update","image_url":"http://xx.cdn","click_url":"http://xxx.play.m.jaeapp.com"}，action为update，代表新增(image_url以及click_url必传)。action=get，代表获得商家设置的活动，image_url以及click_url不用填写。action＝del,代表将活动从资源位撤下
    
    BusinessParams   string `json:"business_params,omitempty"`
    

    // 不用传
    
    Position   string `json:"position,omitempty"`
    

    // 目前必须填0，代表店铺中宝箱资源位
    
    SubBizType   string `json:"sub_biz_type,omitempty"`
    

}
*/

// TopIsvDecorateParam 
type TopIsvDecorateParam struct {

    // 活动id，调用alibaba.interact.activity.register传入的bizid
    BizId   string `json:"biz_id,omitempty"`

    // 目前必须填0，代表店铺
    BizType   string `json:"biz_type,omitempty"`

    // {"action":"update","image_url":"http://xx.cdn","click_url":"http://xxx.play.m.jaeapp.com"}，action为update，代表新增(image_url以及click_url必传)。action=get，代表获得商家设置的活动，image_url以及click_url不用填写。action＝del,代表将活动从资源位撤下
    BusinessParams   string `json:"business_params,omitempty"`

    // 不用传
    Position   string `json:"position,omitempty"`

    // 目前必须填0，代表店铺中宝箱资源位
    SubBizType   string `json:"sub_biz_type,omitempty"`

}
