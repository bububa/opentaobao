package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoCloudprintClientinfoPut 云打印客户端监控信息收集
// cainiao.cloudprint.clientinfo.put
//
// 云打印客户端监控信息收集
func CainiaoCloudprintClientinfoPut(ctx context.Context, clt *core.SDKClient, req *waybill.CainiaoCloudprintClientinfoPutAPIRequest, resp *waybill.CainiaoCloudprintClientinfoPutAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
