package wdk

// Deliveryman 
type Deliveryman struct {
    // 姓名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 编号
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 手机号
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
}
