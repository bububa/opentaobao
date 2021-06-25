package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取营销积木块列表 APIRequest
taobao.ump.mbbs.get

获取营销积木块列表，可以根据类型获取，也可以将该字段设为空，获取所有的
*/
type TaobaoUmpMbbsGetRequest struct {
    model.Params

    // 积木块类型。如果该字段为空表示查出所有类型的<br/>现在有且仅有如下几种：resource,condition,action,target
    type   string 

}

func NewTaobaoUmpMbbsGetRequest() *TaobaoUmpMbbsGetRequest{
    return &TaobaoUmpMbbsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpMbbsGetRequest) GetApiMethodName() string {
    return "taobao.ump.mbbs.get"
}

func (r TaobaoUmpMbbsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpMbbsGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoUmpMbbsGetRequest) GetType() string {
    return r.type
}

