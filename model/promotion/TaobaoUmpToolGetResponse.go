package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询工具 APIResponse
taobao.ump.tool.get

根据工具id获取一个工具对象
*/
type TaobaoUmpToolGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUmpToolGetResponse `json:"taobao_ump_tool_get_response,omitempty"`
}

type TaobaoUmpToolGetResponse struct {

    // 工具信息内容，格式为json，可以通过提供给的sdk里面的MarketingBuilder来处理这个内容
    Content   string `json:"content,omitempty"`

}
