package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaordcaligeniuswarehousereverseeventupdate 销退单事件回传接口
// taobao.rdc.aligenius.warehouse.reverse.event.update
//
// 用于erp回传销退单相关信息到平台
func Taobaordcaligeniuswarehousereverseeventupdate(clt *core.SDKClient, req *logistic.TaobaordcaligeniuswarehousereverseeventupdateAPIRequest, session string) (*logistic.TaobaordcaligeniuswarehousereverseeventupdateAPIResponse, error) {
	var resp logistic.TaobaordcaligeniuswarehousereverseeventupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
