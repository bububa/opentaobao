package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询工具 APIRequest
taobao.ump.tool.get

根据工具id获取一个工具对象
*/
type TaobaoUmpToolGetRequest struct {
    model.Params

    // 工具的id
    toolId   int64 

}

func NewTaobaoUmpToolGetRequest() *TaobaoUmpToolGetRequest{
    return &TaobaoUmpToolGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpToolGetRequest) GetApiMethodName() string {
    return "taobao.ump.tool.get"
}

func (r TaobaoUmpToolGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpToolGetRequest) SetToolId(toolId int64) error {
    r.toolId = toolId
    r.Set("tool_id", toolId)
    return nil
}

func (r TaobaoUmpToolGetRequest) GetToolId() int64 {
    return r.toolId
}

