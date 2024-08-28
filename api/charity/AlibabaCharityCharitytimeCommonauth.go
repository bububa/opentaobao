package charity

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityCharitytimeCommonauth 通用授权
// alibaba.charity.charitytime.commonauth
//
// 三小时和外部账号绑定通用top 返回跳转链接进行绑定
func AlibabaCharityCharitytimeCommonauth(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaCharityCharitytimeCommonauthAPIRequest, resp *charity.AlibabaCharityCharitytimeCommonauthAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
