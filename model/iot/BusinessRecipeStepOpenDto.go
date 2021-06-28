package iot

// BusinessRecipeStepOpenDto 
/* model for simplify = false
type BusinessRecipeStepOpenDto struct {

    // 行业食谱id
    
    BusinessRecipeId   int64 `json:"business_recipe_id,omitempty"`
    

    // 行业食谱步骤id
    
    BusinessRecipeStepId   int64 `json:"business_recipe_step_id,omitempty"`
    

    // 步骤详细描述
    
    Description   string `json:"description,omitempty"`
    

    // 食谱步骤图url
    
    ImageUrl  *struct {
        ImageUrlDto  *ImageUrlDto `json:"image_url_dto,omitempty"`
    } `json:"image_url,omitempty"`
    

    // 食谱步骤指令列表
    
    RecipeStepActionList  struct {
        BusinessRecipeStepActionOpenDto  []BusinessRecipeStepActionOpenDto `json:"business_recipe_step_action_open_dto,omitempty"`
    } `json:"recipe_step_action_list,omitempty"`
    

    // 食谱步骤名
    
    RecipeStepName   string `json:"recipe_step_name,omitempty"`
    

    // 食谱步骤时间
    
    RecipeStepTime   int64 `json:"recipe_step_time,omitempty"`
    

    // 顺序号
    
    Sequence   int64 `json:"sequence,omitempty"`
    

    // 食谱步骤提示
    
    Tips   string `json:"tips,omitempty"`
    

    // 食谱步骤视频
    
    VideoUrl  *struct {
        VideoUrlDto  *VideoUrlDto `json:"video_url_dto,omitempty"`
    } `json:"video_url,omitempty"`
    

}
*/

// BusinessRecipeStepOpenDto 
type BusinessRecipeStepOpenDto struct {

    // 行业食谱id
    BusinessRecipeId   int64 `json:"business_recipe_id,omitempty"`

    // 行业食谱步骤id
    BusinessRecipeStepId   int64 `json:"business_recipe_step_id,omitempty"`

    // 步骤详细描述
    Description   string `json:"description,omitempty"`

    // 食谱步骤图url
    ImageUrl   *ImageUrlDto `json:"image_url,omitempty"`

    // 食谱步骤指令列表
    RecipeStepActionList   []BusinessRecipeStepActionOpenDto `json:"recipe_step_action_list,omitempty"`

    // 食谱步骤名
    RecipeStepName   string `json:"recipe_step_name,omitempty"`

    // 食谱步骤时间
    RecipeStepTime   int64 `json:"recipe_step_time,omitempty"`

    // 顺序号
    Sequence   int64 `json:"sequence,omitempty"`

    // 食谱步骤提示
    Tips   string `json:"tips,omitempty"`

    // 食谱步骤视频
    VideoUrl   *VideoUrlDto `json:"video_url,omitempty"`

}
