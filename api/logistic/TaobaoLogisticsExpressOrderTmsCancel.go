package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticsexpressordertmscancel 服务商上门取退时间取消接口
// taobao.logistics.express.order.tms.cancel
//
// 服务商上门取退时间取消接口
func Taobaologisticsexpressordertmscancel(clt *core.SDKClient, req *logistic.TaobaologisticsexpressordertmscancelAPIRequest, session string) (*logistic.TaobaologisticsexpressordertmscancelAPIResponse, error) {
	var resp logistic.TaobaologisticsexpressordertmscancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
