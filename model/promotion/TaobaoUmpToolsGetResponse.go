package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询工具列表 APIResponse
taobao.ump.tools.get

查询工具列表
*/
type TaobaoUmpToolsGetAPIResponse struct {
    model.CommonResponse
    TaobaoUmpToolsGetResponse
}

type TaobaoUmpToolsGetResponse struct {
    XMLName xml.Name `xml:"ump_tools_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 工具列表，单个内容为json格式，需要通过ump的sdk提供的MarketingBuilder来进行处理
    
    Tools   []string `json:"tools,omitempty" xml:"tools>string,omitempty"`
    
    
}
