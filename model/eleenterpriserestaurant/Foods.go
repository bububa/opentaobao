package eleenterpriserestaurant

// Foods 结构体
type Foods struct {
	// 食物名称
	FoodName string `json:"food_name,omitempty" xml:"food_name,omitempty"`
	// 最近一个月的评价
	RecentRating string `json:"recent_rating,omitempty" xml:"recent_rating,omitempty"`
	// 食物价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 食物图片
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 食物描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 最近一个月的售出份数
	RecentPopularity int64 `json:"recent_popularity,omitempty" xml:"recent_popularity,omitempty"`
	// 食物id
	FoodId int64 `json:"food_id,omitempty" xml:"food_id,omitempty"`
	// 库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
}
