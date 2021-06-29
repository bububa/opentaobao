package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过ids列表获取营销积木块列表 API请求
taobao.ump.mbbs.list.get

通过营销积木id列表来获取营销积木块列表。
*/
type TaobaoUmpMbbsListGetRequest struct {
    model.Params
    // 营销积木块id组成的字符串。
    ids   []int64
}

// 初始化TaobaoUmpMbbsListGetRequest对象
func NewTaobaoUmpMbbsListGetRequest() *TaobaoUmpMbbsListGetRequest{
    return &TaobaoUmpMbbsListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpMbbsListGetRequest) GetApiMethodName() string {
    return "taobao.ump.mbbs.list.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpMbbsListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ids Setter
// 营销积木块id组成的字符串。
func (r *TaobaoUmpMbbsListGetRequest) SetIds(ids []int64) error {
    r.ids = ids
    r.Set("ids", ids)
    return nil
}

// Ids Getter
func (r TaobaoUmpMbbsListGetRequest) GetIds() []int64 {
    return r.ids
}
