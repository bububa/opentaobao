package tanx

// NativeTemplateDto 
/* model for simplify = false
type NativeTemplateDto struct {

    // 样式预览图片url
    
    Preview   string `json:"preview,omitempty"`
    

    // 说明
    
    Description   string `json:"description,omitempty"`
    

    // 模板ID
    
    TmplId   int64 `json:"tmpl_id,omitempty"`
    

    // 区域列表
    
    Areas  struct {
        NativeTemplateAreaDto  []NativeTemplateAreaDto `json:"native_template_area_dto,omitempty"`
    } `json:"areas,omitempty"`
    

    // 模板支持的广告位尺寸
    
    Size   string `json:"size,omitempty"`
    

}
*/

// NativeTemplateDto 
type NativeTemplateDto struct {

    // 样式预览图片url
    Preview   string `json:"preview,omitempty"`

    // 说明
    Description   string `json:"description,omitempty"`

    // 模板ID
    TmplId   int64 `json:"tmpl_id,omitempty"`

    // 区域列表
    Areas   []NativeTemplateAreaDto `json:"areas,omitempty"`

    // 模板支持的广告位尺寸
    Size   string `json:"size,omitempty"`

}
