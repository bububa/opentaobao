package legalcase

// Content 
type Content struct {

    // code值
    
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
    

    // 文本值
    
    Text   string `json:"text,omitempty" xml:"text,omitempty"`
    

    // 二级值集
    
    Children   []Children `json:"children,omitempty" xml:"children,omitempty"`
    

}
