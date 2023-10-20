package iot

import (
	"sync"
)

// RecipeIngredientDto 结构体
type RecipeIngredientDto struct {
	// 参数名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 重量及单位
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
}

var poolRecipeIngredientDto = sync.Pool{
	New: func() any {
		return new(RecipeIngredientDto)
	},
}

// GetRecipeIngredientDto() 从对象池中获取RecipeIngredientDto
func GetRecipeIngredientDto() *RecipeIngredientDto {
	return poolRecipeIngredientDto.Get().(*RecipeIngredientDto)
}

// ReleaseRecipeIngredientDto 释放RecipeIngredientDto
func ReleaseRecipeIngredientDto(v *RecipeIngredientDto) {
	v.Name = ""
	v.Weight = ""
	poolRecipeIngredientDto.Put(v)
}
