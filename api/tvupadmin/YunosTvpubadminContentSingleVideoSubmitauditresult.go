package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentSingleVideoSubmitauditresult 单视频审核提交审核结果
// yunos.tvpubadmin.content.single.video.submitauditresult
//
// 单视频审核提交审核结果
func YunosTvpubadminContentSingleVideoSubmitauditresult(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest, resp *tvupadmin.YunosTvpubadminContentSingleVideoSubmitauditresultAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
