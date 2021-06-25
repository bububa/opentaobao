package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
(新)销量明星删除创意相关接口 APIRequest
taobao.simba.salestar.creative.delete

删除一个创意
*/
type TaobaoSimbaSalestarCreativeDeleteRequest struct {
    model.Params

    // 创意Id
    creativeId   int64 

}

func NewTaobaoSimbaSalestarCreativeDeleteRequest() *TaobaoSimbaSalestarCreativeDeleteRequest{
    return &TaobaoSimbaSalestarCreativeDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaSalestarCreativeDeleteRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.creative.delete"
}

func (r TaobaoSimbaSalestarCreativeDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaSalestarCreativeDeleteRequest) SetCreativeId(creativeId int64) error {
    r.creativeId = creativeId
    r.Set("creative_id", creativeId)
    return nil
}

func (r TaobaoSimbaSalestarCreativeDeleteRequest) GetCreativeId() int64 {
    return r.creativeId
}

