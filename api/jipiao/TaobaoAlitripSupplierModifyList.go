package jipiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jipiao"
)

/* 
【机票代理商订单】改签通知单列表 
taobao.alitrip.supplier.modify.list

提供供应商查询改签通知单列表
*/
func TaobaoAlitripSupplierModifyList(clt *core.SDKClient, req *jipiao.TaobaoAlitripSupplierModifyListAPIRequest, session string) (*jipiao.TaobaoAlitripSupplierModifyListAPIResponse, error) {
    var resp jipiao.TaobaoAlitripSupplierModifyListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
