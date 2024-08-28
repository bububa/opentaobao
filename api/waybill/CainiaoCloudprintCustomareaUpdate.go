package waybill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoCloudprintCustomareaUpdate 自定义区内容更新
// cainiao.cloudprint.customarea.update
//
// 自定义区内容更新
func CainiaoCloudprintCustomareaUpdate(ctx context.Context, clt *core.SDKClient, req *waybill.CainiaoCloudprintCustomareaUpdateAPIRequest, resp *waybill.CainiaoCloudprintCustomareaUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
