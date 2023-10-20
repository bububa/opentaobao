package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaonextonelogisticswarehouseupdate AG退货入仓状态写接口
// taobao.nextone.logistics.warehouse.update
//
// 商家上传退货入仓状态给ag
func Taobaonextonelogisticswarehouseupdate(clt *core.SDKClient, req *logistic.TaobaonextonelogisticswarehouseupdateAPIRequest, session string) (*logistic.TaobaonextonelogisticswarehouseupdateAPIResponse, error) {
	var resp logistic.TaobaonextonelogisticswarehouseupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
