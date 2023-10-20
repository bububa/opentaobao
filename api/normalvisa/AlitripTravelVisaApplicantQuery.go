package normalvisa

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// Alitriptravelvisaapplicantquery 签证申请人查询接口
// alitrip.travel.visa.applicant.query
//
// 签证申请人查询接口，商家可根据条件查询申请人id，用于签证办理
func Alitriptravelvisaapplicantquery(clt *core.SDKClient, req *normalvisa.AlitriptravelvisaapplicantqueryAPIRequest, session string) (*normalvisa.AlitriptravelvisaapplicantqueryAPIResponse, error) {
	var resp normalvisa.AlitriptravelvisaapplicantqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
