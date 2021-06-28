package waybill

// RenderConfig 
/* model for simplify = false
type RenderConfig struct {

    // 打印方向：normal-正常 reverse-翻转(旋转180°)
    
    Orientation   string `json:"orientation,omitempty"`
    

    // 下联logo
    
    NeedBottomLogo   bool `json:"need_bottom_logo,omitempty"`
    

    // 中间logo
    
    NeedMiddleLogo   bool `json:"need_middle_logo,omitempty"`
    

    // 上联logo
    
    NeedTopLogo   bool `json:"need_top_logo,omitempty"`
    

}
*/

// RenderConfig 
type RenderConfig struct {

    // 打印方向：normal-正常 reverse-翻转(旋转180°)
    Orientation   string `json:"orientation,omitempty"`

    // 下联logo
    NeedBottomLogo   bool `json:"need_bottom_logo,omitempty"`

    // 中间logo
    NeedMiddleLogo   bool `json:"need_middle_logo,omitempty"`

    // 上联logo
    NeedTopLogo   bool `json:"need_top_logo,omitempty"`

}
