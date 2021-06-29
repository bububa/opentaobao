package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取营销积木块 API请求
taobao.ump.mbb.getbyid

根据积木块id获取营销积木块。
*/
type TaobaoUmpMbbGetbyidRequest struct {
    model.Params
    // 积木块的id
    id   int64
}

// 初始化TaobaoUmpMbbGetbyidRequest对象
func NewTaobaoUmpMbbGetbyidRequest() *TaobaoUmpMbbGetbyidRequest{
    return &TaobaoUmpMbbGetbyidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpMbbGetbyidRequest) GetApiMethodName() string {
    return "taobao.ump.mbb.getbyid"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpMbbGetbyidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 积木块的id
func (r *TaobaoUmpMbbGetbyidRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TaobaoUmpMbbGetbyidRequest) GetId() int64 {
    return r.id
}
