package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentVideoSubmitauditresult 迎客松提交视频审核结果
// yunos.tvpubadmin.content.video.submitauditresult
//
// 迎客松提交视频审核结果
func YunosTvpubadminContentVideoSubmitauditresult(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentVideoSubmitauditresultAPIRequest, resp *tvupadmin.YunosTvpubadminContentVideoSubmitauditresultAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
