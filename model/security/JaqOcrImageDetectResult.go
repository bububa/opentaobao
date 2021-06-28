package security

// JaqOcrImageDetectResult 
/* model for simplify = false
type JaqOcrImageDetectResult struct {

    // 字符串列表，内容是图像中文字的主要段落内容（按照概率输出最多5组概率最大的组合）
    
    Texts  struct {
        String  []string `json:"string,omitempty"`
    } `json:"texts,omitempty"`
    

}
*/

// JaqOcrImageDetectResult 
type JaqOcrImageDetectResult struct {

    // 字符串列表，内容是图像中文字的主要段落内容（按照概率输出最多5组概率最大的组合）
    Texts   []string `json:"texts,omitempty"`

}
