package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供货商查询订单数据接口 APIRequest
cainiao.cntec.supplier.order.service

提供给供货商查询订单信息的接口，返回给供货商的订单数据已经是脱敏精简后的，比如订单ID用户ID已经用md5加密，用户昵称已经脱敏，商品信息本身是供货商提供的。
数据查询的范围只和供货商的身份有关系，比如大润发的用户只能查询大润发的订单，而且会校验身份和颁发的appkey之间的关系，并且目前对接的只有一个供货商
*/
type CainiaoCntecSupplierOrderServiceRequest struct {
    model.Params

    // 系统自动生成
    queryConditioin   *SupplierOrderQueryDto 

}

func NewCainiaoCntecSupplierOrderServiceRequest() *CainiaoCntecSupplierOrderServiceRequest{
    return &CainiaoCntecSupplierOrderServiceRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCntecSupplierOrderServiceRequest) GetApiMethodName() string {
    return "cainiao.cntec.supplier.order.service"
}

func (r CainiaoCntecSupplierOrderServiceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCntecSupplierOrderServiceRequest) SetQueryConditioin(queryConditioin *SupplierOrderQueryDto) error {
    r.queryConditioin = queryConditioin
    r.Set("query_conditioin", queryConditioin)
    return nil
}

func (r CainiaoCntecSupplierOrderServiceRequest) GetQueryConditioin() *SupplierOrderQueryDto {
    return r.queryConditioin
}

