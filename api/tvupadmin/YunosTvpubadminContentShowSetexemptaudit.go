package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentShowSetexemptaudit 迎客松节目设置免审开关
// yunos.tvpubadmin.content.show.setexemptaudit
//
// 迎客松节目设置免审开关
func YunosTvpubadminContentShowSetexemptaudit(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentShowSetexemptauditAPIRequest, resp *tvupadmin.YunosTvpubadminContentShowSetexemptauditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
