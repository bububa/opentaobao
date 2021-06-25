package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询工具列表 APIRequest
taobao.ump.tools.get

查询工具列表
*/
type TaobaoUmpToolsGetRequest struct {
    model.Params

    // 工具编码
    toolCode   string 

}

func NewTaobaoUmpToolsGetRequest() *TaobaoUmpToolsGetRequest{
    return &TaobaoUmpToolsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpToolsGetRequest) GetApiMethodName() string {
    return "taobao.ump.tools.get"
}

func (r TaobaoUmpToolsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpToolsGetRequest) SetToolCode(toolCode string) error {
    r.toolCode = toolCode
    r.Set("tool_code", toolCode)
    return nil
}

func (r TaobaoUmpToolsGetRequest) GetToolCode() string {
    return r.toolCode
}

