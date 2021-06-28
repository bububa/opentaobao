package scbp

// CrowdTagView 
type CrowdTagView struct {

    // 标签中文名
    
    TagName   string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
    

    // 标签id(潜在访问偏好和潜在采购意向返回的是类目id，店铺老客和优选人群返回的是字符串)
    
    TagId   string `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
    

    // 溢价百分比
    
    Discount   int64 `json:"discount,omitempty" xml:"discount,omitempty"`
    

    // 最近7天效果数据
    
    Effect   *Effect7d `json:"effect,omitempty" xml:"effect,omitempty"`
    

    // 人群类型 PRI_VISIT(访问过本店铺产品的买家) PRI_INQUIRY(给本店铺发起过询盘的买家) PUB_VISIT_CATE(潜在访问偏好) PUB_INQUIRY_CATE(潜在采购意向) PUBLIC_REGION(国家) PUB_HIGH_MOQ(高MOQ偏好买家) PUB_HIGH_RETURN(高回访买家) PUB_HIGH_POTENTIAL_ORDER(潜在交易买家) PUB_HIGH_POTENTIAL_AB(高活跃买家)
    
    CrowdType   string `json:"crowd_type,omitempty" xml:"crowd_type,omitempty"`
    

}
