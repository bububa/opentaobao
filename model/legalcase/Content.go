package legalcase

// Content 
/* model for simplify = false
type Content struct {

    // code值
    
    Value   string `json:"value,omitempty"`
    

    // 文本值
    
    Text   string `json:"text,omitempty"`
    

    // 二级值集
    
    Children  struct {
        Children  []Children `json:"children,omitempty"`
    } `json:"children,omitempty"`
    

}
*/

// Content 
type Content struct {

    // code值
    Value   string `json:"value,omitempty"`

    // 文本值
    Text   string `json:"text,omitempty"`

    // 二级值集
    Children   []Children `json:"children,omitempty"`

}
