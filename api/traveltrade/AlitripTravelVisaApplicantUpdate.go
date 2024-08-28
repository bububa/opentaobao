package traveltrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelVisaApplicantUpdate 飞猪度假-普通签证-申请人进度推进接口
// alitrip.travel.visa.applicant.update
//
// 普通签证订单的申请人进度推进接口，用于商家代用户填写申请人基本信息 或 推进单个申请人的签证进度。
func AlitripTravelVisaApplicantUpdate(ctx context.Context, clt *core.SDKClient, req *traveltrade.AlitripTravelVisaApplicantUpdateAPIRequest, resp *traveltrade.AlitripTravelVisaApplicantUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
