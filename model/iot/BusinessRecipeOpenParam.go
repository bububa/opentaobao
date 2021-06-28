package iot

// BusinessRecipeOpenParam 
/* model for simplify = false
type BusinessRecipeOpenParam struct {

    // 菜谱描述
    
    Description   string `json:"description,omitempty"`
    

    // 品类Id
    
    DevTypeId   int64 `json:"dev_type_id,omitempty"`
    

    // 食谱id
    
    RecipeId   int64 `json:"recipe_id,omitempty"`
    

    // 开放账号id
    
    OpenAccountId   string `json:"open_account_id,omitempty"`
    

    // 食谱图片
    
    RecipeImage  *struct {
        ImageUrlParam  *ImageUrlParam `json:"image_url_param,omitempty"`
    } `json:"recipe_image,omitempty"`
    

    // 重量及单位
    
    RecipeIngredientList  struct {
        RecipeIngredientParam  []RecipeIngredientParam `json:"recipe_ingredient_param,omitempty"`
    } `json:"recipe_ingredient_list,omitempty"`
    

    // 食谱中文名
    
    RecipeNameCn   string `json:"recipe_name_cn,omitempty"`
    

    // 食谱类型 0视频菜谱 1图文菜谱
    
    RecipeType   int64 `json:"recipe_type,omitempty"`
    

    // 食谱视频
    
    RecipeVideo  *struct {
        VideoUrlParam  *VideoUrlParam `json:"video_url_param,omitempty"`
    } `json:"recipe_video,omitempty"`
    

    // 标签列表
    
    TagIdList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"tag_id_list,omitempty"`
    

    // 菜谱提示信息
    
    Tips   string `json:"tips,omitempty"`
    

}
*/

// BusinessRecipeOpenParam 
type BusinessRecipeOpenParam struct {

    // 菜谱描述
    Description   string `json:"description,omitempty"`

    // 品类Id
    DevTypeId   int64 `json:"dev_type_id,omitempty"`

    // 食谱id
    RecipeId   int64 `json:"recipe_id,omitempty"`

    // 开放账号id
    OpenAccountId   string `json:"open_account_id,omitempty"`

    // 食谱图片
    RecipeImage   *ImageUrlParam `json:"recipe_image,omitempty"`

    // 重量及单位
    RecipeIngredientList   []RecipeIngredientParam `json:"recipe_ingredient_list,omitempty"`

    // 食谱中文名
    RecipeNameCn   string `json:"recipe_name_cn,omitempty"`

    // 食谱类型 0视频菜谱 1图文菜谱
    RecipeType   int64 `json:"recipe_type,omitempty"`

    // 食谱视频
    RecipeVideo   *VideoUrlParam `json:"recipe_video,omitempty"`

    // 标签列表
    TagIdList   []int64 `json:"tag_id_list,omitempty"`

    // 菜谱提示信息
    Tips   string `json:"tips,omitempty"`

}
