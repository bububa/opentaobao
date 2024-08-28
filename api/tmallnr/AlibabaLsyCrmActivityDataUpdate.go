package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmActivityDataUpdate 私域导购数据回流接口
// alibaba.lsy.crm.activity.data.update
//
// 私域导购数据回流接口
func AlibabaLsyCrmActivityDataUpdate(ctx context.Context, clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmActivityDataUpdateAPIRequest, resp *tmallnr.AlibabaLsyCrmActivityDataUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
