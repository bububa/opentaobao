package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据营销积木块代码获取积木块 APIRequest
taobao.ump.mbb.getbycode

根据营销积木块代码获取积木块。接口返回该代码最新版本的积木块。如果要查询某个非最新版本的积木块，可以使用积木块id查询接口。
*/
type TaobaoUmpMbbGetbycodeRequest struct {
    model.Params

    // 营销积木块code
    code   string 

}

func NewTaobaoUmpMbbGetbycodeRequest() *TaobaoUmpMbbGetbycodeRequest{
    return &TaobaoUmpMbbGetbycodeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpMbbGetbycodeRequest) GetApiMethodName() string {
    return "taobao.ump.mbb.getbycode"
}

func (r TaobaoUmpMbbGetbycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpMbbGetbycodeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r TaobaoUmpMbbGetbycodeRequest) GetCode() string {
    return r.code
}

