package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeThreedimensionSync 二手房3D户型信息同步
// alibaba.alihouse.existinghome.threedimension.sync
//
// 二手房3D户型信息同步
func AlibabaAlihouseExistinghomeThreedimensionSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeThreedimensionSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
