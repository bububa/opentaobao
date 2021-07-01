package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelVisaApplicantQueryAPIRequest
签证申请人查询接口 API请求
alitrip.travel.visa.applicant.query

签证申请人查询接口，商家可根据条件查询申请人id，用于签证办理 */
type AlitripTravelVisaApplicantQueryAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *QueryApplicantParam
}

// New
