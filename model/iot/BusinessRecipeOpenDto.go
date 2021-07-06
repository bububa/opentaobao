package iot

// BusinessRecipeOpenDto 结构体
type BusinessRecipeOpenDto struct {
	// 食谱食材参数列表
	RecipeIngredientList []RecipeIngredientDto `json:"recipe_ingredient_list,omitempty" xml:"recipe_ingredient_list>recipe_ingredient_dto,omitempty"`
	// 食谱步骤列表
	RecipeStepList []BusinessRecipeStepOpenDto `json:"recipe_step_list,omitempty" xml:"recipe_step_list>business_recipe_step_open_dto,omitempty"`
	// 标签列表
	TagList []ContentTagDto `json:"tag_list,omitempty" xml:"tag_list>content_tag_dto,omitempty"`
	// 菜谱描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 食谱中文名
	RecipeNameCn string `json:"recipe_name_cn,omitempty" xml:"recipe_name_cn,omitempty"`
	// 食谱提示信息
	Tips string `json:"tips,omitempty" xml:"tips,omitempty"`
	// 食谱类型 0视频菜谱 1图文菜谱
	RecipeType string `json:"recipe_type,omitempty" xml:"recipe_type,omitempty"`
	// 行业食谱id
	BusinessRecipeId int64 `json:"business_recipe_id,omitempty" xml:"business_recipe_id,omitempty"`
	// 菜谱功能类型 1 普通食谱 2 智能食谱
	FunctionType int64 `json:"function_type,omitempty" xml:"function_type,omitempty"`
	// 食谱时间，单位秒
	RecipeTime int64 `json:"recipe_time,omitempty" xml:"recipe_time,omitempty"`
	// 食谱视频
	RecipeVideo *VideoUrlDto `json:"recipe_video,omitempty" xml:"recipe_video,omitempty"`
	// 食谱图片
	RecipeImage *ImageUrlDto `json:"recipe_image,omitempty" xml:"recipe_image,omitempty"`
}
