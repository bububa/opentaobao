package cloudgame

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCgameLiteplayAvatarRecordReport Avatar形象保存地址回调
// alibaba.cgame.liteplay.avatar.record.report
//
// 新氢玩, 围观互动Avatar捏脸, 形象地址保存回调
func AlibabaCgameLiteplayAvatarRecordReport(ctx context.Context, clt *core.SDKClient, req *cloudgame.AlibabaCgameLiteplayAvatarRecordReportAPIRequest, resp *cloudgame.AlibabaCgameLiteplayAvatarRecordReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
