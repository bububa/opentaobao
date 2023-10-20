package baoxian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baoxian"
)

// Alipaybaoxianclaimreturngoodsstatusupdate 更新理赔单退货货物状态
// alipay.baoxian.claim.returngoodsstatus.update
//
// 更新理赔单退货货物状态
func Alipaybaoxianclaimreturngoodsstatusupdate(clt *core.SDKClient, req *baoxian.AlipaybaoxianclaimreturngoodsstatusupdateAPIRequest, session string) (*baoxian.AlipaybaoxianclaimreturngoodsstatusupdateAPIResponse, error) {
	var resp baoxian.AlipaybaoxianclaimreturngoodsstatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
