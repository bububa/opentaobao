package baoxian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baoxian"
)

// AlipayBaoxianClaimSurveyConclusionSubmit 保险退货服务商勘察结论提交接口
// alipay.baoxian.claim.survey.conclusion.submit
//
// 保险退货服务商提交勘察结论
func AlipayBaoxianClaimSurveyConclusionSubmit(clt *core.SDKClient, req *baoxian.AlipayBaoxianClaimSurveyConclusionSubmitAPIRequest, resp *baoxian.AlipayBaoxianClaimSurveyConclusionSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
