package xhotelonlineorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelonlineorder"
)

/* 
客人PMS账单信息查询 
taobao.xhotel.pms.guestbill.get.vtwo

从pms获取客人账单信息
*/
func TaobaoXhotelPmsGuestbillGetVtwo(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelPmsGuestbillGetVtwoRequest, session string) (*xhotelonlineorder.TaobaoXhotelPmsGuestbillGetVtwoAPIResponse, error) {
    var resp xhotelonlineorder.TaobaoXhotelPmsGuestbillGetVtwoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
