package cainiaohandover

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverParcelQuery 获取交接单小包信息
// cainiao.global.handover.parcel.query
//
// 提供给ISV通过该接口查询小包信息
func CainiaoGlobalHandoverParcelQuery(ctx context.Context, clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverParcelQueryAPIRequest, resp *cainiaohandover.CainiaoGlobalHandoverParcelQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
