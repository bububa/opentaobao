package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过ids列表获取营销积木块列表 APIRequest
taobao.ump.mbbs.list.get

通过营销积木id列表来获取营销积木块列表。
*/
type TaobaoUmpMbbsListGetRequest struct {
    model.Params

    // 营销积木块id组成的字符串。
    ids   []int64 

}

func NewTaobaoUmpMbbsListGetRequest() *TaobaoUmpMbbsListGetRequest{
    return &TaobaoUmpMbbsListGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpMbbsListGetRequest) GetApiMethodName() string {
    return "taobao.ump.mbbs.list.get"
}

func (r TaobaoUmpMbbsListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpMbbsListGetRequest) SetIds(ids []int64) error {
    r.ids = ids
    r.Set("ids", ids)
    return nil
}

func (r TaobaoUmpMbbsListGetRequest) GetIds() []int64 {
    return r.ids
}

