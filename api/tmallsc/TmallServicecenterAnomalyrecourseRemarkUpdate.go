package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseRemarkUpdate 天猫服务平台一键求助单服务商备注更新接口
// tmall.servicecenter.anomalyrecourse.remark.update
//
// 一键求助服务商可以回传备注
func TmallServicecenterAnomalyrecourseRemarkUpdate(ctx context.Context, clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseRemarkUpdateAPIRequest, resp *tmallsc.TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
