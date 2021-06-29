package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询工具 API请求
taobao.ump.tool.get

根据工具id获取一个工具对象
*/
type TaobaoUmpToolGetRequest struct {
    model.Params
    // 工具的id
    _toolId   int64
}

// 初始化TaobaoUmpToolGetRequest对象
func NewTaobaoUmpToolGetRequest() *TaobaoUmpToolGetRequest{
    return &TaobaoUmpToolGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpToolGetRequest) GetApiMethodName() string {
    return "taobao.ump.tool.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpToolGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ToolId Setter
// 工具的id
func (r *TaobaoUmpToolGetRequest) SetToolId(_toolId int64) error {
    r._toolId = _toolId
    r.Set("tool_id", _toolId)
    return nil
}

// ToolId Getter
func (r TaobaoUmpToolGetRequest) GetToolId() int64 {
    return r._toolId
}
