package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaordcaligeniuswarehouseresendlogisticsmsgpost 补发单物流信息回传
// taobao.rdc.aligenius.warehouse.resend.logistics.msg.post
//
// 补发单erp物流信息回传平台
func Taobaordcaligeniuswarehouseresendlogisticsmsgpost(clt *core.SDKClient, req *logistic.TaobaordcaligeniuswarehouseresendlogisticsmsgpostAPIRequest, session string) (*logistic.TaobaordcaligeniuswarehouseresendlogisticsmsgpostAPIResponse, error) {
	var resp logistic.TaobaordcaligeniuswarehouseresendlogisticsmsgpostAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
