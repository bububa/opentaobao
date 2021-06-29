package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取营销积木块列表 API请求
taobao.ump.mbbs.get

获取营销积木块列表，可以根据类型获取，也可以将该字段设为空，获取所有的
*/
type TaobaoUmpMbbsGetRequest struct {
    model.Params
    // 积木块类型。如果该字段为空表示查出所有类型的<br/>现在有且仅有如下几种：resource,condition,action,target
    _type   string
}

// 初始化TaobaoUmpMbbsGetRequest对象
func NewTaobaoUmpMbbsGetRequest() *TaobaoUmpMbbsGetRequest{
    return &TaobaoUmpMbbsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpMbbsGetRequest) GetApiMethodName() string {
    return "taobao.ump.mbbs.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpMbbsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 积木块类型。如果该字段为空表示查出所有类型的<br/>现在有且仅有如下几种：resource,condition,action,target
func (r *TaobaoUmpMbbsGetRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoUmpMbbsGetRequest) GetType() string {
    return r._type
}
