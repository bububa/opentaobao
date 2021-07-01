package wlb

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlb"
)

/* 
家装发货接口 
taobao.wlb.order.jzwithins.consign

为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发带安装服务商的发货接口
*/
func TaobaoWlbOrderJzwithinsConsign(clt *core.SDKClient, req *wlb.TaobaoWlbOrderJzwithinsConsignAPIRequest, session string) (*wlb.TaobaoWlbOrderJzwithinsConsignAPIResponse, error) {
    var resp wlb.TaobaoWlbOrderJzwithinsConsignAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
