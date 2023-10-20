package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverParcelQuery 获取交接单小包信息
// cainiao.global.handover.parcel.query
//
// 提供给ISV通过该接口查询小包信息
func CainiaoGlobalHandoverParcelQuery(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverParcelQueryAPIRequest, resp *cainiaohandover.CainiaoGlobalHandoverParcelQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
