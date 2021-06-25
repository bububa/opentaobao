package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判断用户是否收藏某个店铺 APIRequest
alibaba.interact.open.isattention

判断用户是否收藏某个店铺
*/
type AlibabaInteractOpenIsattentionRequest struct {
    model.Params

    // 1
    param0   int64 

}

func NewAlibabaInteractOpenIsattentionRequest() *AlibabaInteractOpenIsattentionRequest{
    return &AlibabaInteractOpenIsattentionRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractOpenIsattentionRequest) GetApiMethodName() string {
    return "alibaba.interact.open.isattention"
}

func (r AlibabaInteractOpenIsattentionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractOpenIsattentionRequest) SetParam0(param0 int64) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaInteractOpenIsattentionRequest) GetParam0() int64 {
    return r.param0
}

