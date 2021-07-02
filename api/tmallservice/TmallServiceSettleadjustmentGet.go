package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServiceSettleadjustmentGet 查询结算调整单单条记录
// tmall.service.settleadjustment.get
//
// 提供给服务商通过结算调整单id获取结算调整单信息
func TmallServiceSettleadjustmentGet(clt *core.SDKClient, req *tmallservice.TmallServiceSettleadjustmentGetAPIRequest, session string) (*tmallservice.TmallServiceSettleadjustmentGetAPIResponse, error) {
	var resp tmallservice.TmallServiceSettleadjustmentGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
