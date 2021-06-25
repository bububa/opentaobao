package iot

// BusinessRecipeOpenDto 
type BusinessRecipeOpenDto struct {

    // 行业食谱id
    BusinessRecipeId   int64 `json:"business_recipe_id,omitempty"`

    // 菜谱描述
    Description   string `json:"description,omitempty"`

    // 菜谱功能类型 1 普通食谱 2 智能食谱
    FunctionType   int64 `json:"function_type,omitempty"`

    // 食谱食材参数列表
    RecipeIngredientList   []RecipeIngredientDto `json:"recipe_ingredient_list,omitempty"`

    // 食谱中文名
    RecipeNameCn   string `json:"recipe_name_cn,omitempty"`

    // 食谱步骤列表
    RecipeStepList   []BusinessRecipeStepOpenDto `json:"recipe_step_list,omitempty"`

    // 食谱时间，单位秒
    RecipeTime   int64 `json:"recipe_time,omitempty"`

    // 食谱视频
    RecipeVideo   *VideoUrlDto `json:"recipe_video,omitempty"`

    // 标签列表
    TagList   []ContentTagDto `json:"tag_list,omitempty"`

    // 食谱提示信息
    Tips   string `json:"tips,omitempty"`

    // 食谱图片
    RecipeImage   *ImageUrlDto `json:"recipe_image,omitempty"`

    // 食谱类型 0视频菜谱 1图文菜谱
    RecipeType   string `json:"recipe_type,omitempty"`

}
