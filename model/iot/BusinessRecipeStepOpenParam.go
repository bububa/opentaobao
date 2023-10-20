package iot

import (
	"sync"
)

// BusinessRecipeStepOpenParam 结构体
type BusinessRecipeStepOpenParam struct {
	// 步骤指令列表
	RecipeStepActionList []BusinessRecipeStepActionOpenParam `json:"recipe_step_action_list,omitempty" xml:"recipe_step_action_list>business_recipe_step_action_open_param,omitempty"`
	// 食谱步骤描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 开放账号Id
	OpenAccountId string `json:"open_account_id,omitempty" xml:"open_account_id,omitempty"`
	// 食谱步骤中文名
	RecipeStepNameCn string `json:"recipe_step_name_cn,omitempty" xml:"recipe_step_name_cn,omitempty"`
	// 食谱步骤提示
	Tips string `json:"tips,omitempty" xml:"tips,omitempty"`
	// 行业食谱id
	BusinessRecipeId int64 `json:"business_recipe_id,omitempty" xml:"business_recipe_id,omitempty"`
	// 行业食谱步骤id
	BusinessRecipeStepId int64 `json:"business_recipe_step_id,omitempty" xml:"business_recipe_step_id,omitempty"`
	// 食谱步骤图片
	ImageUrl *ImageUrlParam `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 食谱步骤时间
	RecipeStepTime int64 `json:"recipe_step_time,omitempty" xml:"recipe_step_time,omitempty"`
	// 食谱步骤顺序号
	Sequence int64 `json:"sequence,omitempty" xml:"sequence,omitempty"`
	// 食谱步骤视频
	VideoUrl *VideoUrlParam `json:"video_url,omitempty" xml:"video_url,omitempty"`
	// 指令标识：0不支持指令，1支持指令
	ActionFlag int64 `json:"action_flag,omitempty" xml:"action_flag,omitempty"`
}

var poolBusinessRecipeStepOpenParam = sync.Pool{
	New: func() any {
		return new(BusinessRecipeStepOpenParam)
	},
}

// GetBusinessRecipeStepOpenParam() 从对象池中获取BusinessRecipeStepOpenParam
func GetBusinessRecipeStepOpenParam() *BusinessRecipeStepOpenParam {
	return poolBusinessRecipeStepOpenParam.Get().(*BusinessRecipeStepOpenParam)
}

// ReleaseBusinessRecipeStepOpenParam 释放BusinessRecipeStepOpenParam
func ReleaseBusinessRecipeStepOpenParam(v *BusinessRecipeStepOpenParam) {
	v.RecipeStepActionList = v.RecipeStepActionList[:0]
	v.Description = ""
	v.OpenAccountId = ""
	v.RecipeStepNameCn = ""
	v.Tips = ""
	v.BusinessRecipeId = 0
	v.BusinessRecipeStepId = 0
	v.ImageUrl = nil
	v.RecipeStepTime = 0
	v.Sequence = 0
	v.VideoUrl = nil
	v.ActionFlag = 0
	poolBusinessRecipeStepOpenParam.Put(v)
}
