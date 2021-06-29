package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询货架和位置数据 APIRequest
alibaba.interact.retail.queryshelflocation

查询货架和位置数据
*/
type AlibabaInteractRetailQueryshelflocationRequest struct {
    model.Params

    // 门店code
    param0   string 

}

func NewAlibabaInteractRetailQueryshelflocationRequest() *AlibabaInteractRetailQueryshelflocationRequest{
    return &AlibabaInteractRetailQueryshelflocationRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractRetailQueryshelflocationRequest) GetApiMethodName() string {
    return "alibaba.interact.retail.queryshelflocation"
}

func (r AlibabaInteractRetailQueryshelflocationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractRetailQueryshelflocationRequest) SetParam0(param0 string) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaInteractRetailQueryshelflocationRequest) GetParam0() string {
    return r.param0
}

