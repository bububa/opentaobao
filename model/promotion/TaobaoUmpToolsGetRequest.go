package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询工具列表 API请求
taobao.ump.tools.get

查询工具列表
*/
type TaobaoUmpToolsGetRequest struct {
    model.Params
    // 工具编码
    _toolCode   string
}

// 初始化TaobaoUmpToolsGetRequest对象
func NewTaobaoUmpToolsGetRequest() *TaobaoUmpToolsGetRequest{
    return &TaobaoUmpToolsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpToolsGetRequest) GetApiMethodName() string {
    return "taobao.ump.tools.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpToolsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ToolCode Setter
// 工具编码
func (r *TaobaoUmpToolsGetRequest) SetToolCode(_toolCode string) error {
    r._toolCode = _toolCode
    r.Set("tool_code", _toolCode)
    return nil
}

// ToolCode Getter
func (r TaobaoUmpToolsGetRequest) GetToolCode() string {
    return r._toolCode
}
