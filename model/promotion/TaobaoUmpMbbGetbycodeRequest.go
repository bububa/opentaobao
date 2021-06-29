package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据营销积木块代码获取积木块 API请求
taobao.ump.mbb.getbycode

根据营销积木块代码获取积木块。接口返回该代码最新版本的积木块。如果要查询某个非最新版本的积木块，可以使用积木块id查询接口。
*/
type TaobaoUmpMbbGetbycodeRequest struct {
    model.Params
    // 营销积木块code
    _code   string
}

// 初始化TaobaoUmpMbbGetbycodeRequest对象
func NewTaobaoUmpMbbGetbycodeRequest() *TaobaoUmpMbbGetbycodeRequest{
    return &TaobaoUmpMbbGetbycodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpMbbGetbycodeRequest) GetApiMethodName() string {
    return "taobao.ump.mbb.getbycode"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpMbbGetbycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 营销积木块code
func (r *TaobaoUmpMbbGetbycodeRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r TaobaoUmpMbbGetbycodeRequest) GetCode() string {
    return r._code
}
