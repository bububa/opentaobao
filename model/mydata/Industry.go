package mydata

// Industry 
type Industry struct {

    // 行业描述
    
    IndustryDesc   string `json:"industry_desc,omitempty" xml:"industry_desc,omitempty"`
    

    // 行业ID
    
    IndustryId   int64 `json:"industry_id,omitempty" xml:"industry_id,omitempty"`
    

    // 是否主营行业
    
    MainCategory   bool `json:"main_category,omitempty" xml:"main_category,omitempty"`
    

}
