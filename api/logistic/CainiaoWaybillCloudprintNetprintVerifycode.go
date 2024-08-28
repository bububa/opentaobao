package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// CainiaoWaybillCloudprintNetprintVerifycode 打印验证码
// cainiao.waybill.cloudprint.netprint.verifycode
//
// 打印获取验证码
func CainiaoWaybillCloudprintNetprintVerifycode(ctx context.Context, clt *core.SDKClient, req *logistic.CainiaoWaybillCloudprintNetprintVerifycodeAPIRequest, resp *logistic.CainiaoWaybillCloudprintNetprintVerifycodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
