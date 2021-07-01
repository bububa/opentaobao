package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
工单维度虚拟中间号绑定 
tmall.servicecenter.workcard.virtualphone.bind

服务供应链洗护服务ERP项目中，客服呼叫消费者的功能。
叫消费者的手机号虚拟号码给到客服。
*/
func TmallServicecenterWorkcardVirtualphoneBind(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardVirtualphoneBindAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardVirtualphoneBindAPIResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardVirtualphoneBindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
