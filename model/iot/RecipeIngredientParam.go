package iot

import (
	"sync"
)

// RecipeIngredientParam 结构体
type RecipeIngredientParam struct {
	// 参数名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 重量及单位
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
}

var poolRecipeIngredientParam = sync.Pool{
	New: func() any {
		return new(RecipeIngredientParam)
	},
}

// GetRecipeIngredientParam() 从对象池中获取RecipeIngredientParam
func GetRecipeIngredientParam() *RecipeIngredientParam {
	return poolRecipeIngredientParam.Get().(*RecipeIngredientParam)
}

// ReleaseRecipeIngredientParam 释放RecipeIngredientParam
func ReleaseRecipeIngredientParam(v *RecipeIngredientParam) {
	v.Name = ""
	v.Weight = ""
	poolRecipeIngredientParam.Put(v)
}
