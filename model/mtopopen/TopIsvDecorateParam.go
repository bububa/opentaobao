package mtopopen

// TopIsvDecorateParam 结构体
type TopIsvDecorateParam struct {
	// 活动id，调用alibaba.interact.activity.register传入的bizid
	BizId string `json:"biz_id,omitempty" xml:"biz_id,omitempty"`
	// 目前必须填0，代表店铺
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// {&#34;action&#34;:&#34;update&#34;,&#34;image_url&#34;:&#34;http://xx.cdn&#34;,&#34;click_url&#34;:&#34;http://xxx.play.m.jaeapp.com&#34;}，action为update，代表新增(image_url以及click_url必传)。action=get，代表获得商家设置的活动，image_url以及click_url不用填写。action＝del,代表将活动从资源位撤下
	BusinessParams string `json:"business_params,omitempty" xml:"business_params,omitempty"`
	// 不用传
	Position string `json:"position,omitempty" xml:"position,omitempty"`
	// 目前必须填0，代表店铺中宝箱资源位
	SubBizType string `json:"sub_biz_type,omitempty" xml:"sub_biz_type,omitempty"`
}
