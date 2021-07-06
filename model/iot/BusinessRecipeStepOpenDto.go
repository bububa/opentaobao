package iot

// BusinessRecipeStepOpenDto 结构体
type BusinessRecipeStepOpenDto struct {
	// 食谱步骤指令列表
	RecipeStepActionList []BusinessRecipeStepActionOpenDto `json:"recipe_step_action_list,omitempty" xml:"recipe_step_action_list>business_recipe_step_action_open_dto,omitempty"`
	// 步骤详细描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 食谱步骤名
	RecipeStepName string `json:"recipe_step_name,omitempty" xml:"recipe_step_name,omitempty"`
	// 食谱步骤提示
	Tips string `json:"tips,omitempty" xml:"tips,omitempty"`
	// 行业食谱id
	BusinessRecipeId int64 `json:"business_recipe_id,omitempty" xml:"business_recipe_id,omitempty"`
	// 行业食谱步骤id
	BusinessRecipeStepId int64 `json:"business_recipe_step_id,omitempty" xml:"business_recipe_step_id,omitempty"`
	// 食谱步骤图url
	ImageUrl *ImageUrlDto `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 食谱步骤时间
	RecipeStepTime int64 `json:"recipe_step_time,omitempty" xml:"recipe_step_time,omitempty"`
	// 顺序号
	Sequence int64 `json:"sequence,omitempty" xml:"sequence,omitempty"`
	// 食谱步骤视频
	VideoUrl *VideoUrlDto `json:"video_url,omitempty" xml:"video_url,omitempty"`
}
