package product

// SpuTemplateDO 
type SpuTemplateDO struct {

    // 产品关键属性，内容为属性ID(PID)的列表，注意关键属性可以在类目上不存在。不存在的PID，默认为输入，没有子属性。属性名称在prop_name_str中取
    
    KeyProperties   []int64 `json:"key_properties,omitempty" xml:"key_properties>int64,omitempty"`
    

    // 产品绑定属性，内容为属性ID(PID)的列表,绑定属性肯定在类目上有，对应属性的类目特征，子属性请根据PID到类目上去取
    
    AffectProperties   []int64 `json:"affect_properties,omitempty" xml:"affect_properties>int64,omitempty"`
    

    // 过滤属性，内容有属性ID(PID)列表，很重要的属性，filter_properties包含的属性，必须填写
    
    FilterProperties   []int64 `json:"filter_properties,omitempty" xml:"filter_properties>int64,omitempty"`
    

    // 类目ID
    
    CategoryId   int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
    

    // 模板ID，发布产品，必须放到Product中
    
    TemplateId   int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
    

    // 品类ID，和类目ID类似
    
    CommodityId   int64 `json:"commodity_id,omitempty" xml:"commodity_id,omitempty"`
    

    // 属性名称扁平化结构，只保证不在类目上的CP有值
    
    PropNameStr   string `json:"prop_name_str,omitempty" xml:"prop_name_str,omitempty"`
    

    // 预留
    
    PropFeatures   string `json:"prop_features,omitempty" xml:"prop_features,omitempty"`
    

}
