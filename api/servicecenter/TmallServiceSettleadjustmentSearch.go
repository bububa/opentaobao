package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
服务商15分钟获取数据 
tmall.service.settleadjustment.search

天猫服务平台，按修改时间，时间间隔在15中内（包含15分钟），获取调整单数据
*/
func TmallServiceSettleadjustmentSearch(clt *core.SDKClient, req *servicecenter.TmallServiceSettleadjustmentSearchAPIRequest, session string) (*servicecenter.TmallServiceSettleadjustmentSearchAPIResponse, error) {
    var resp servicecenter.TmallServiceSettleadjustmentSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
