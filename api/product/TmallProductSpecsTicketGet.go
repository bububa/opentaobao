package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
产品规格审核信息获取接口 
tmall.product.specs.ticket.get

批量根据specId查询产品规格审核信息包括产品规格状态，申请人，拒绝原因等
*/
func TmallProductSpecsTicketGet(clt *core.SDKClient, req *product.TmallProductSpecsTicketGetAPIRequest, session string) (*product.TmallProductSpecsTicketGetAPIResponse, error) {
    var resp product.TmallProductSpecsTicketGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
