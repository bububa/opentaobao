package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品规格审核信息获取接口 API请求
tmall.product.specs.ticket.get

批量根据specId查询产品规格审核信息包括产品规格状态，申请人，拒绝原因等
*/
type TmallProductSpecsTicketGetRequest struct {
    model.Params
    // 产品规格ID，多个用逗号分隔
    _specIds   string
}

// 初始化TmallProductSpecsTicketGetRequest对象
func NewTmallProductSpecsTicketGetRequest() *TmallProductSpecsTicketGetRequest{
    return &TmallProductSpecsTicketGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallProductSpecsTicketGetRequest) GetApiMethodName() string {
    return "tmall.product.specs.ticket.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallProductSpecsTicketGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SpecIds Setter
// 产品规格ID，多个用逗号分隔
func (r *TmallProductSpecsTicketGetRequest) SetSpecIds(_specIds string) error {
    r._specIds = _specIds
    r.Set("spec_ids", _specIds)
    return nil
}

// SpecIds Getter
func (r TmallProductSpecsTicketGetRequest) GetSpecIds() string {
    return r._specIds
}
