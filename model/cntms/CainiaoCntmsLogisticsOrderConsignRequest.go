package cntms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟配商家仓库发货 APIRequest
cainiao.cntms.logistics.order.consign

商家包装打印面单结束后，通知菜鸟包裹要发货
*/
type CainiaoCntmsLogisticsOrderConsignRequest struct {
    model.Params

    // 配送发货信息
    content   *CnTmsLogisticsOrderConsignContent 

}

func NewCainiaoCntmsLogisticsOrderConsignRequest() *CainiaoCntmsLogisticsOrderConsignRequest{
    return &CainiaoCntmsLogisticsOrderConsignRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCntmsLogisticsOrderConsignRequest) GetApiMethodName() string {
    return "cainiao.cntms.logistics.order.consign"
}

func (r CainiaoCntmsLogisticsOrderConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCntmsLogisticsOrderConsignRequest) SetContent(content *CnTmsLogisticsOrderConsignContent) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r CainiaoCntmsLogisticsOrderConsignRequest) GetContent() *CnTmsLogisticsOrderConsignContent {
    return r.content
}

