package eleenterpriserestaurant

// Activity 结构体
type Activity struct {
	// 活动名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
}
