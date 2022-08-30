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
	// 属性内容类型,1:"单选",2:"多选",3:"字符",4:"长字符",5:"数字",6:"日期"
	PropertyContentType int64 `json:"property_content_type,omitempty" xml:"property_content_type,omitempty"`
	// 属性类型
	CategoryAttrType int64 `json:"category_attr_type,omitempty" xml:"category_attr_type,omitempty"`
	// 属性规则对象
	PropertyRule *PropertyRuleDto `json:"property_rule,omitempty" xml:"property_rule,omitempty"`
	// 属性等级
	PropertyGrade int64 `json:"property_grade,omitempty" xml:"property_grade,omitempty"`
	// 属性ID
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
	// 内容类型,1:"常规",2:"游戏帐号",3:"游戏密码",4:"QQ号",5:"单价数量",6:"买家服务器",7:"开始发货时间",8:"终止发货时间",9:"原价",10:"联系手机",11:"账户类型",12:"发货服务器",13:"角色名称",14:"备用名称",15:"第三名称"  ,16:"确认帐号",17:"解锁密码",18:"面额",19:"充值类型",20:"回收折扣",21:"帐号绑定",22:"多选"
	ContentType int64 `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// 是否可见
	Visible bool `json:"visible,omitempty" xml:"visible,omitempty"`
}
