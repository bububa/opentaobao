package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
修改结算调整单 
tmall.service.settleadjustment.modify

提供给服务商在对结算有异议时，发起结算调整单。
通过说明调整单ID，调整费用值，调整原因进行结算调整单修改。
*/
func TmallServiceSettleadjustmentModify(clt *core.SDKClient, req *user.TmallServiceSettleadjustmentModifyRequest, session string) (*user.TmallServiceSettleadjustmentModifyAPIResponse, error) {
    var resp user.TmallServiceSettleadjustmentModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
