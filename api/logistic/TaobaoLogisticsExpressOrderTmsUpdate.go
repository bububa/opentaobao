package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticsexpressordertmsupdate 服务商修改上门取退时间接口
// taobao.logistics.express.order.tms.update
//
// 服务商修改上门取退时间接口
func Taobaologisticsexpressordertmsupdate(clt *core.SDKClient, req *logistic.TaobaologisticsexpressordertmsupdateAPIRequest, session string) (*logistic.TaobaologisticsexpressordertmsupdateAPIResponse, error) {
	var resp logistic.TaobaologisticsexpressordertmsupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
