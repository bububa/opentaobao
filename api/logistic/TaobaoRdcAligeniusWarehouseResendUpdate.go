package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaordcaligeniuswarehouseresendupdate 补发单状态回传
// taobao.rdc.aligenius.warehouse.resend.update
//
// 补发单状态回传接口
func Taobaordcaligeniuswarehouseresendupdate(clt *core.SDKClient, req *logistic.TaobaordcaligeniuswarehouseresendupdateAPIRequest, session string) (*logistic.TaobaordcaligeniuswarehouseresendupdateAPIResponse, error) {
	var resp logistic.TaobaordcaligeniuswarehouseresendupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
