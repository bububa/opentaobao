package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
创建代扣计划 
alibaba.idle.pay.plan.create

闲鱼平台代扣能力：
1、用户和闲鱼签约代扣协议 服务商通过直付通产品挂载成为闲鱼二级商户 来完成用户和服务商的结算
2、创建代扣计划
*/
func AlibabaIdlePayPlanCreate(clt *core.SDKClient, req *idle.AlibabaIdlePayPlanCreateRequest, session string) (*idle.AlibabaIdlePayPlanCreateAPIResponse, error) {
    var resp idle.AlibabaIdlePayPlanCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
