package wlbimports

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wlbimports"
)

/* 
集货商家pdf和云打印面单获取，pdf需要配置白名单 
taobao.wlb.crossborder.waybill.get

【TOF】欧洲供应商PDF格式电子面单渲染下发
 需求链接：https://aone.alibaba-inc.com/req/21210808
*/
func TaobaoWlbCrossborderWaybillGet(clt *core.SDKClient, req *wlbimports.TaobaoWlbCrossborderWaybillGetRequest, session string) (*wlbimports.TaobaoWlbCrossborderWaybillGetAPIResponse, error) {
    var resp wlbimports.TaobaoWlbCrossborderWaybillGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
