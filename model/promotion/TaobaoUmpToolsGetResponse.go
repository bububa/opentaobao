package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询工具列表 APIResponse
taobao.ump.tools.get

查询工具列表
*/
type TaobaoUmpToolsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUmpToolsGetResponse `json:"taobao_ump_tools_get_response,omitempty"`
}

type TaobaoUmpToolsGetResponse struct {

    // 工具列表，单个内容为json格式，需要通过ump的sdk提供的MarketingBuilder来进行处理
    Tools   []String `json:"tools,omitempty"`

}
