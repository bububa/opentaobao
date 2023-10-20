package iot

import (
	"sync"
)

// BusinessRecipeOpenParam 结构体
type BusinessRecipeOpenParam struct {
	// 重量及单位
	RecipeIngredientList []RecipeIngredientParam `json:"recipe_ingredient_list,omitempty" xml:"recipe_ingredient_list>recipe_ingredient_param,omitempty"`
	// 标签列表
	TagIdList []int64 `json:"tag_id_list,omitempty" xml:"tag_id_list>int64,omitempty"`
	// 菜谱描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 开放账号id
	OpenAccountId string `json:"open_account_id,omitempty" xml:"open_account_id,omitempty"`
	// 食谱中文名
	RecipeNameCn string `json:"recipe_name_cn,omitempty" xml:"recipe_name_cn,omitempty"`
	// 菜谱提示信息
	Tips string `json:"tips,omitempty" xml:"tips,omitempty"`
	// 品类Id
	DevTypeId int64 `json:"dev_type_id,omitempty" xml:"dev_type_id,omitempty"`
	// 食谱id
	RecipeId int64 `json:"recipe_id,omitempty" xml:"recipe_id,omitempty"`
	// 食谱图片
	RecipeImage *ImageUrlParam `json:"recipe_image,omitempty" xml:"recipe_image,omitempty"`
	// 食谱类型 0视频菜谱 1图文菜谱
	RecipeType int64 `json:"recipe_type,omitempty" xml:"recipe_type,omitempty"`
	// 食谱视频
	RecipeVideo *VideoUrlParam `json:"recipe_video,omitempty" xml:"recipe_video,omitempty"`
}

var poolBusinessRecipeOpenParam = sync.Pool{
	New: func() any {
		return new(BusinessRecipeOpenParam)
	},
}

// GetBusinessRecipeOpenParam() 从对象池中获取BusinessRecipeOpenParam
func GetBusinessRecipeOpenParam() *BusinessRecipeOpenParam {
	return poolBusinessRecipeOpenParam.Get().(*BusinessRecipeOpenParam)
}

// ReleaseBusinessRecipeOpenParam 释放BusinessRecipeOpenParam
func ReleaseBusinessRecipeOpenParam(v *BusinessRecipeOpenParam) {
	v.RecipeIngredientList = v.RecipeIngredientList[:0]
	v.TagIdList = v.TagIdList[:0]
	v.Description = ""
	v.OpenAccountId = ""
	v.RecipeNameCn = ""
	v.Tips = ""
	v.DevTypeId = 0
	v.RecipeId = 0
	v.RecipeImage = nil
	v.RecipeType = 0
	v.RecipeVideo = nil
	poolBusinessRecipeOpenParam.Put(v)
}
