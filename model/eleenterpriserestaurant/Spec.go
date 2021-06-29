package eleenterpriserestaurant

// Spec 
type Spec struct {
    // 姓名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 值
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
}
