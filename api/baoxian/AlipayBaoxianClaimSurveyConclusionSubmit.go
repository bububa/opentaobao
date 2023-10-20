package baoxian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baoxian"
)

// Alipaybaoxianclaimsurveyconclusionsubmit 保险退货服务商勘察结论提交接口
// alipay.baoxian.claim.survey.conclusion.submit
//
// 保险退货服务商提交勘察结论
func Alipaybaoxianclaimsurveyconclusionsubmit(clt *core.SDKClient, req *baoxian.AlipaybaoxianclaimsurveyconclusionsubmitAPIRequest, session string) (*baoxian.AlipaybaoxianclaimsurveyconclusionsubmitAPIResponse, error) {
	var resp baoxian.AlipaybaoxianclaimsurveyconclusionsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
