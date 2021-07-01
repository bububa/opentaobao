package iot

// RecipeIngredientDto 结构体
type RecipeIngredientDto struct {
	// 参数名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 重量及单位
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
}
