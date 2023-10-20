package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaowlbimportthreeplofflineconsign 3PL直邮线下发货
// taobao.wlb.import.threepl.offline.consign
//
// 菜鸟认证直邮线下发货
func Taobaowlbimportthreeplofflineconsign(clt *core.SDKClient, req *logistic.TaobaowlbimportthreeplofflineconsignAPIRequest, session string) (*logistic.TaobaowlbimportthreeplofflineconsignAPIResponse, error) {
	var resp logistic.TaobaowlbimportthreeplofflineconsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
