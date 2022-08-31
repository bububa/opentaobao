package tmalltrend

// StyleBasicInfoBo 结构体
type StyleBasicInfoBo struct {
	// 款式同步目的，枚举，INSERT(&#34;新增&#34;), UPDATE(&#34;更新&#34;), OFFLINE(&#34;下线&#34;);
	SyncPurpose string `json:"sync_purpose,omitempty" xml:"sync_purpose,omitempty"`
	// 款式下线原因，枚举，当上面同步目的字段为DELETE时此参数必传，IP_COPYRIGHT(&#34;版权原因&#34;), DESIGN_FLAW(&#34;设计缺陷&#34;), OTHER(&#34;其他原因&#34;);
	OfflineReason string `json:"offline_reason,omitempty" xml:"offline_reason,omitempty"`
	// 版权归属方，集团安全约束，此参数必填
	CopyrightOwner string `json:"copyright_owner,omitempty" xml:"copyright_owner,omitempty"`
	// 款式模型源文件URL
	ModelSourceUrl string `json:"model_source_url,omitempty" xml:"model_source_url,omitempty"`
	// 款式编号，业务唯一
	StyleSerialNumber string `json:"style_serial_number,omitempty" xml:"style_serial_number,omitempty"`
	// 图案元素，多个用英文分号分割组成字符串传递
	Pattern string `json:"pattern,omitempty" xml:"pattern,omitempty"`
	// 款式名称
	StyleName string `json:"style_name,omitempty" xml:"style_name,omitempty"`
	// 工艺技法，多个用英文分号分割
	Craftsmanship string `json:"craftsmanship,omitempty" xml:"craftsmanship,omitempty"`
	// 主料类别，枚举，KNITTING(&#34;针织&#34;), WOVEN(&#34;梭织&#34;), WOOLEN(&#34;毛织&#34;), NON_WOVEN_FABRIC(&#34;牛仔&#34;), LEATHER(&#34;皮草&#34;)
	MajorMaterialCate string `json:"major_material_cate,omitempty" xml:"major_material_cate,omitempty"`
	// 风格类型，枚举，YOUNG_LADY_CLOTHING(&#34;少淑女装&#34;), JIAN_OU_ZHONG_SHU(&#34;简欧中淑&#34;), MATURE_AND_ELEGANT(&#34;成熟优雅&#34;), COTTON_LINEN_STYLE(&#34;棉麻风&#34;), STREET_FASHION(&#34;棉麻风&#34;), GUO_CHAO(&#34;国朝&#34;), BUSINESS_CASUAL(&#34;商务休闲&#34;), FASHION_FUNCTION(&#34;时尚机能&#34;), SPORTS_AND_LEISURE(&#34;运动休闲&#34;), OUTDOOR_SPORTS(&#34;运动休闲&#34;), JAPANESE_AND_KOREAN_STYLE(&#34;日韩风格&#34;), COLLEGE_STYLE(&#34;学院风&#34;), CHINESE_STYLE(&#34;中国风&#34;), FOLK_CUSTOM(&#34;民族风&#34;), BOHEMIA_STYLE(&#34;波西米亚&#34;), MIDDLE_EASTERN_STYLE(&#34;中东风&#34;), WORKPIECE_STYLE(&#34;工装风&#34;);
	StyleType string `json:"style_type,omitempty" xml:"style_type,omitempty"`
	// 款式主图链接
	MainPicSourceUrl string `json:"main_pic_source_url,omitempty" xml:"main_pic_source_url,omitempty"`
	// 款式细节图，多个链接地址用英文分号分割，最多10张细节图
	DetailPicSourceUrl string `json:"detail_pic_source_url,omitempty" xml:"detail_pic_source_url,omitempty"`
	// 廓形分类
	ContourCate string `json:"contour_cate,omitempty" xml:"contour_cate,omitempty"`
	// 辅料
	MinorMaterial string `json:"minor_material,omitempty" xml:"minor_material,omitempty"`
	// 核心属性项，英文分号分割，最多3个，必须按照给定备选项设置，否则强校验会不通过
	KeyProperties string `json:"key_properties,omitempty" xml:"key_properties,omitempty"`
	// 款式价格，单位元
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 授权证明，图片链接，当版权方不是凌迪时
	AuthorizeProofSourceUrl string `json:"authorize_proof_source_url,omitempty" xml:"authorize_proof_source_url,omitempty"`
	// 偏好人群，枚举，LUXURIOUS_CROWED(&#34;奢华人群&#34;), FASHION_CROWED(&#34;奢华人群&#34;), MAIN_STREAM_FASHION_CROWED(&#34;主流时尚者&#34;), QUALITY_PRAGMATIC_CROWED(&#34;品质实用者&#34;), SHOPAHOLIC(&#34;购物狂&#34;), ECONOMICAL_PRAGMATIC_CROWED(&#34;平价实惠者&#34;)
	CrowedType string `json:"crowed_type,omitempty" xml:"crowed_type,omitempty"`
	// 装饰细节，多个用英文分号分割
	DecorationDetail string `json:"decoration_detail,omitempty" xml:"decoration_detail,omitempty"`
	// 主料成分
	MajorMaterialComponent string `json:"major_material_component,omitempty" xml:"major_material_component,omitempty"`
	// 款式视频地址链接
	VideoSourceUrl string `json:"video_source_url,omitempty" xml:"video_source_url,omitempty"`
	// 下线操作时说明详细下线原因
	OfflineReasonDetail string `json:"offline_reason_detail,omitempty" xml:"offline_reason_detail,omitempty"`
	// 服装模型信息
	Clothings string `json:"clothings,omitempty" xml:"clothings,omitempty"`
	// 人台模型信息
	Modelling string `json:"modelling,omitempty" xml:"modelling,omitempty"`
	// 部件模型信息
	Components string `json:"components,omitempty" xml:"components,omitempty"`
	// 面料信息
	Fabrics string `json:"fabrics,omitempty" xml:"fabrics,omitempty"`
	// 图案信息
	Decals string `json:"decals,omitempty" xml:"decals,omitempty"`
	// 字体图案信息
	Fonts string `json:"fonts,omitempty" xml:"fonts,omitempty"`
	// 辅料信息
	Accessories string `json:"accessories,omitempty" xml:"accessories,omitempty"`
	// 场景信息
	Scenes string `json:"scenes,omitempty" xml:"scenes,omitempty"`
	// 灯光、控制信息
	Lights string `json:"lights,omitempty" xml:"lights,omitempty"`
	// 版本号
	Version string `json:"version,omitempty" xml:"version,omitempty"`
	// 是否拥有IP版权，true--有，false--无
	IpCopyright bool `json:"ip_copyright,omitempty" xml:"ip_copyright,omitempty"`
}
