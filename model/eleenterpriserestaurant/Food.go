package eleenterpriserestaurant

// Food 结构体
type Food struct {
	// 食物信息
	Specfoods []Specfood `json:"specfoods,omitempty" xml:"specfoods>specfood,omitempty"`
	// 活动
	Activity *Activity `json:"activity,omitempty" xml:"activity,omitempty"`
	// 餐品id
	FoodId string `json:"food_id,omitempty" xml:"food_id,omitempty"`
	// 评分
	Rating string `json:"rating,omitempty" xml:"rating,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 规格信息, 例如: [{“name”: “规格”, “values”: [“不辣”, “中辣”, “大辣” ]}, …]
	Specifications []Specification `json:"specifications,omitempty" xml:"specifications>specification,omitempty"`
	// 概况说明, 例: “2评价 月售6份”
	Tips string `json:"tips,omitempty" xml:"tips,omitempty"`
	// 属性信息, 次要的规格信息, 不决定价格, 例如: [{“name”: “甜度”, “values”: [“加糖”, “不加糖”]}, …]
	Attrs []Attr `json:"attrs,omitempty" xml:"attrs>attr,omitempty"`
	// 月售份数
	MonthSales int64 `json:"month_sales,omitempty" xml:"month_sales,omitempty"`
	// 图片地址
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 最少起购分数, 默认: 1
	MinPurchase int64 `json:"min_purchase,omitempty" xml:"min_purchase,omitempty"`
	// 好评率
	SatisfyRate int64 `json:"satisfy_rate,omitempty" xml:"satisfy_rate,omitempty"`
	// 是否必点菜品
	IsEssential bool `json:"is_essential,omitempty" xml:"is_essential,omitempty"`
}
