package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
(新)销量明星删除创意相关接口 API请求
taobao.simba.salestar.creative.delete

删除一个创意
*/
type TaobaoSimbaSalestarCreativeDeleteRequest struct {
    model.Params
    // 创意Id
    _creativeId   int64
}

// 初始化TaobaoSimbaSalestarCreativeDeleteRequest对象
func NewTaobaoSimbaSalestarCreativeDeleteRequest() *TaobaoSimbaSalestarCreativeDeleteRequest{
    return &TaobaoSimbaSalestarCreativeDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarCreativeDeleteRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.creative.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarCreativeDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreativeId Setter
// 创意Id
func (r *TaobaoSimbaSalestarCreativeDeleteRequest) SetCreativeId(_creativeId int64) error {
    r._creativeId = _creativeId
    r.Set("creative_id", _creativeId)
    return nil
}

// CreativeId Getter
func (r TaobaoSimbaSalestarCreativeDeleteRequest) GetCreativeId() int64 {
    return r._creativeId
}
