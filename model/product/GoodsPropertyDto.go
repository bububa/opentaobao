package product

// GoodsPropertyDto 结构体
type GoodsPropertyDto struct {
	// 属性可选项
	Options []GoodsPropertyOptionDto `json:"options,omitempty" xml:"options>goods_property_option_dto,omitempty"`
	// 默认值
	DefaultValue string `json:"default_value,omitempty" xml:"default_value,omitempty"`
	// 商品属性单位
	AttrUnit string `json:"attr_unit,omitempty" xml:"attr_unit,omitempty"`
	// 属性名称
	PropertyName string `json:"property_name,omitempty" xml:"property_name,omitempty"`
	// 占位符
	Placeholder string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// 属性内容类型,1:&#34;单选&#34;,2:&#34;多选&#34;,3:&#34;字符&#34;,4:&#34;长字符&#34;,5:&#34;数字&#34;,6:&#34;日期&#34;
	PropertyContentType int64 `json:"property_content_type,omitempty" xml:"property_content_type,omitempty"`
	// 属性类型
	CategoryAttrType int64 `json:"category_attr_type,omitempty" xml:"category_attr_type,omitempty"`
	// 属性规则对象
	PropertyRule *PropertyRuleDto `json:"property_rule,omitempty" xml:"property_rule,omitempty"`
	// 属性等级
	PropertyGrade int64 `json:"property_grade,omitempty" xml:"property_grade,omitempty"`
	// 属性ID
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
	// 内容类型,1:&#34;常规&#34;,2:&#34;游戏帐号&#34;,3:&#34;游戏密码&#34;,4:&#34;QQ号&#34;,5:&#34;单价数量&#34;,6:&#34;买家服务器&#34;,7:&#34;开始发货时间&#34;,8:&#34;终止发货时间&#34;,9:&#34;原价&#34;,10:&#34;联系手机&#34;,11:&#34;账户类型&#34;,12:&#34;发货服务器&#34;,13:&#34;角色名称&#34;,14:&#34;备用名称&#34;,15:&#34;第三名称&#34;  ,16:&#34;确认帐号&#34;,17:&#34;解锁密码&#34;,18:&#34;面额&#34;,19:&#34;充值类型&#34;,20:&#34;回收折扣&#34;,21:&#34;帐号绑定&#34;,22:&#34;多选&#34;
	ContentType int64 `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// 是否可见
	Visible bool `json:"visible,omitempty" xml:"visible,omitempty"`
}
