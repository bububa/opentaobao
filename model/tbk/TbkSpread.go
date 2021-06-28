package tbk

// TbkSpread 
type TbkSpread struct {

    // 传播形式, 目前只支持短链接
    
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    

    // 调用错误信息；由于是批量接口，请重点关注每条请求返回的结果，如果非OK，则说明该结果对应的content不正常，请酌情处理;
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

}
