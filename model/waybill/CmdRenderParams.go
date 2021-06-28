package waybill

// CmdRenderParams 
/* model for simplify = false
type CmdRenderParams struct {

    // 需要打印的文档，包括模板地址、打印数据
    
    Document  *struct {
        RenderDocument  *RenderDocument `json:"render_document,omitempty"`
    } `json:"document,omitempty"`
    

    // 打印机名称
    
    PrinterName   string `json:"printer_name,omitempty"`
    

    // 客户端ID
    
    ClientId   string `json:"client_id,omitempty"`
    

    // 客户端类型：weixin/alipay/native
    
    ClientType   string `json:"client_type,omitempty"`
    

    // 打印配置
    
    Config  *struct {
        RenderConfig  *RenderConfig `json:"render_config,omitempty"`
    } `json:"config,omitempty"`
    

}
*/

// CmdRenderParams 
type CmdRenderParams struct {

    // 需要打印的文档，包括模板地址、打印数据
    Document   *RenderDocument `json:"document,omitempty"`

    // 打印机名称
    PrinterName   string `json:"printer_name,omitempty"`

    // 客户端ID
    ClientId   string `json:"client_id,omitempty"`

    // 客户端类型：weixin/alipay/native
    ClientType   string `json:"client_type,omitempty"`

    // 打印配置
    Config   *RenderConfig `json:"config,omitempty"`

}
