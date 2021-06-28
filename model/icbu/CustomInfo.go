package icbu

// CustomInfo 
/* model for simplify = false
type CustomInfo struct {

    // 定制内容
    
    CustomContents  struct {
        CustomContent  []CustomContent `json:"custom_content,omitempty"`
    } `json:"custom_contents,omitempty"`
    

}
*/

// CustomInfo 
type CustomInfo struct {

    // 定制内容
    CustomContents   []CustomContent `json:"custom_contents,omitempty"`

}
