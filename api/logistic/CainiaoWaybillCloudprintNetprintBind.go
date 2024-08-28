package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// CainiaoWaybillCloudprintNetprintBind 网络打印机绑定
// cainiao.waybill.cloudprint.netprint.bind
//
// 绑定打印机接口
func CainiaoWaybillCloudprintNetprintBind(ctx context.Context, clt *core.SDKClient, req *logistic.CainiaoWaybillCloudprintNetprintBindAPIRequest, resp *logistic.CainiaoWaybillCloudprintNetprintBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
