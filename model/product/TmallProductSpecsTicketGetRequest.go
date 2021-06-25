package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品规格审核信息获取接口 APIRequest
tmall.product.specs.ticket.get

批量根据specId查询产品规格审核信息包括产品规格状态，申请人，拒绝原因等
*/
type TmallProductSpecsTicketGetRequest struct {
    model.Params

    // 产品规格ID，多个用逗号分隔
    specIds   string 

}

func NewTmallProductSpecsTicketGetRequest() *TmallProductSpecsTicketGetRequest{
    return &TmallProductSpecsTicketGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallProductSpecsTicketGetRequest) GetApiMethodName() string {
    return "tmall.product.specs.ticket.get"
}

func (r TmallProductSpecsTicketGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallProductSpecsTicketGetRequest) SetSpecIds(specIds string) error {
    r.specIds = specIds
    r.Set("spec_ids", specIds)
    return nil
}

func (r TmallProductSpecsTicketGetRequest) GetSpecIds() string {
    return r.specIds
}

