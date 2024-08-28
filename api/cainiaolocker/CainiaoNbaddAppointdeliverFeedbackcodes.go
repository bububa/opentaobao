package cainiaolocker

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// CainiaoNbaddAppointdeliverFeedbackcodes 服务质量反馈编码列表
// cainiao.nbadd.appointdeliver.feedbackcodes
//
// 服务质量反馈编码列表，建议获取数据后缓存在本地，定时刷新即可
func CainiaoNbaddAppointdeliverFeedbackcodes(ctx context.Context, clt *core.SDKClient, req *cainiaolocker.CainiaoNbaddAppointdeliverFeedbackcodesAPIRequest, resp *cainiaolocker.CainiaoNbaddAppointdeliverFeedbackcodesAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
