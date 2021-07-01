package tanx

// CreativeParamDto 结构体
type CreativeParamDto struct {
	// dsp系统中的创意id
	CreativeId string `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
	// 创意尺寸，长高中间用小写英文x
	AdboardSize string `json:"adboard_size,omitempty" xml:"adboard_size,omitempty"`
	// 广告类目
	AdboardType string `json:"adboard_type,omitempty" xml:"adboard_type,omitempty"`
	// 敏感词类目
	SensitiveType string `json:"sensitive_type,omitempty" xml:"sensitive_type,omitempty"`
	// 创意代码。creative_package_format为5, dis_type=2时，对应的值为标准json，支持的字段为：adwords,clickUrl,imgUrl,landingType,price,promoprice,sell,title如：{&quot;clickUrl&quot;:&quot;http://click.mz.simba.taobao.com/&quot;,&quot;title&quot;:&quot;老板电器旗舰店&mdash;钜惠风暴席卷月！&quot;,&quot;promoprice&quot;:&quot;188.88&quot;,&quot;sell&quot;:&quot;56&quot;, &quot;adwords&quot;:&quot;精选店铺，优惠无止尽！&quot;}
	AdboardData string `json:"adboard_data,omitempty" xml:"adboard_data,omitempty"`
	// 目标地址
	DestinationUrl string `json:"destination_url,omitempty" xml:"destination_url,omitempty"`
	// 创意封装类型：1 Htmlsnippet(pc网页),2 vast-nonlinear（视频暂停）,5 无线创意, 8 移动网页
	CreativePackageFormat int64 `json:"creative_package_format,omitempty" xml:"creative_package_format,omitempty"`
	// 广告主id,多值使用逗号分隔
	AdvertiserIds string `json:"advertiser_ids,omitempty" xml:"advertiser_ids,omitempty"`
	// 创意类型。1. mraid, 2. native, 3.H5，4.富媒体
	DisType int64 `json:"dis_type,omitempty" xml:"dis_type,omitempty"`
	// 模板Id
	TemplateId string `json:"template_id,omitempty" xml:"template_id,omitempty"`
}
