package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServiceSettleadjustmentOperate 天猫服务调整单操作
// tmall.service.settleadjustment.operate
//
// 提供给服务商在对结算有异议时，发起结算调整单。
// 通过说明调整单ID，调整费用值，调整原因进行结算调整单修改。
func TmallServiceSettleadjustmentOperate(clt *core.SDKClient, req *tmallservice.TmallServiceSettleadjustmentOperateAPIRequest, session string) (*tmallservice.TmallServiceSettleadjustmentOperateAPIResponse, error) {
	var resp tmallservice.TmallServiceSettleadjustmentOperateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
