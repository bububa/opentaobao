package iot

// RecipeIngredientParam 结构体
type RecipeIngredientParam struct {
	// 参数名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 重量及单位
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
}
