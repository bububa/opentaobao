package normalvisa

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// AlitripTravelVisaApplicantQuery 签证申请人查询接口
// alitrip.travel.visa.applicant.query
//
// 签证申请人查询接口，商家可根据条件查询申请人id，用于签证办理
func AlitripTravelVisaApplicantQuery(clt *core.SDKClient, req *normalvisa.AlitripTravelVisaApplicantQueryAPIRequest, session string) (*normalvisa.AlitripTravelVisaApplicantQueryAPIResponse, error) {
	var resp normalvisa.AlitripTravelVisaApplicantQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
