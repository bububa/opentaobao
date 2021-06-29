package vaccin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/vaccin"
)

/* 
疫苗接种完成(带支付宝提醒) 
alibaba.health.vaccin.notice.order.complete

用户到店完成接种,ISV感知通知阿里健康完成接种,并通知用户!
*/
func AlibabaHealthVaccinNoticeOrderComplete(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeOrderCompleteRequest, session string) (*vaccin.AlibabaHealthVaccinNoticeOrderCompleteAPIResponse, error) {
    var resp vaccin.AlibabaHealthVaccinNoticeOrderCompleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
