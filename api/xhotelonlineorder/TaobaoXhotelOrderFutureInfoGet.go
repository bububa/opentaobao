package xhotelonlineorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelonlineorder"
)

/* 
获取(查询)订单变更信息 
taobao.xhotel.order.future.info.get

支持操作类型 1.在线开发票请求 3.在线选房请求 4.自助checkIn请求 13.扫脸入住身份信息请求 10.房态信息查询请求 103.通用任务取消指令
*/
func TaobaoXhotelOrderFutureInfoGet(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelOrderFutureInfoGetAPIRequest, session string) (*xhotelonlineorder.TaobaoXhotelOrderFutureInfoGetAPIResponse, error) {
    var resp xhotelonlineorder.TaobaoXhotelOrderFutureInfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
