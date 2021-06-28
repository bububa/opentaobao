package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询工具 APIResponse
taobao.ump.tool.get

根据工具id获取一个工具对象
*/
type TaobaoUmpToolGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ump_tool_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 工具信息内容，格式为json，可以通过提供给的sdk里面的MarketingBuilder来处理这个内容
    
    Content   string `json:"content,omitempty" xml:"