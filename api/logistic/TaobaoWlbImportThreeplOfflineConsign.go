package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoWlbImportThreeplOfflineConsign 3PL直邮线下发货
// taobao.wlb.import.threepl.offline.consign
//
// 菜鸟认证直邮线下发货
func TaobaoWlbImportThreeplOfflineConsign(clt *core.SDKClient, req *logistic.TaobaoWlbImportThreeplOfflineConsignAPIRequest, resp *logistic.TaobaoWlbImportThreeplOfflineConsignAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
