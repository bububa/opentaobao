package cainiaohandover

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverCloudprintGet 获取面单云打印数据
// cainiao.global.handover.cloudprint.get
//
// 提供给ISV通过该接口获取面单云打印数据
func CainiaoGlobalHandoverCloudprintGet(ctx context.Context, clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverCloudprintGetAPIRequest, resp *cainiaohandover.CainiaoGlobalHandoverCloudprintGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
