package tmalltrend

// StyleBasicInfoBo 结构体
type StyleBasicInfoBo struct {
	// 款式同步目的，枚举，INSERT("新增"), UPDATE("更新"), OFFLINE("下线");
	SyncPurpose string `json:"sync_purpose,omitempty" xml:"sync_purpose,omitempty"`
	// 款式下线原因，枚举，当上面同步目的字段为DELETE时此参数必传，IP_COPYRIGHT("版权原因"), DESIGN_FLAW("设计缺陷"), OTHER("其他原因");
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
	// 主料类别，枚举，KNITTING("针织"), WOVEN("梭织"), WOOLEN("毛织"), NON_WOVEN_FABRIC("牛仔"), LEATHER("皮草")
	MajorMaterialCate string `json:"major_material_cate,omitempty" xml:"major_material_cate,omitempty"`
	// 风格类型，枚举，YOUNG_LADY_CLOTHING("少淑女装"), JIAN_OU_ZHONG_SHU("简欧中淑"), MATURE_AND_ELEGANT("成熟优雅"), COTTON_LINEN_STYLE("棉麻风"), STREET_FASHION("棉麻风"), GUO_CHAO("国朝"), BUSINESS_CASUAL("商务休闲"), FASHION_FUNCTION("时尚机能"), SPORTS_AND_LEISURE("运动休闲"), OUTDOOR_SPORTS("运动休闲"), JAPANESE_AND_KOREAN_STYLE("日韩风格"), COLLEGE_STYLE("学院风"), CHINESE_STYLE("中国风"), FOLK_CUSTOM("民族风"), BOHEMIA_STYLE("波西米亚"), MIDDLE_EASTERN_STYLE("中东风"), WORKPIECE_STYLE("工装风");
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
	// 偏好人群，枚举，LUXURIOUS_CROWED("奢华人群"), FASHION_CROWED("奢华人群"), MAIN_STREAM_FASHION_CROWED("主流时尚者"), QUALITY_PRAGMATIC_CROWED("品质实用者"), SHOPAHOLIC("购物狂"), ECONOMICAL_PRAGMATIC_CROWED("平价实惠者")
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
