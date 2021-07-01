package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxQualificationAddAPIRequest
提交资质接口 API请求
taobao.tanx.qualification.add

dsp客户提交客户资质和行业资质 */
type TaobaoTanxQualificationAddAPIRequest struct {
	model.Params
	// dsp客户新增资质dto
	_qualifications []Qualification
	// dsp用户memberId
	_memberId int64
	// dsp验证的token
	_token string
	// 签名时间，1970年到现在的秒
	_signTime int64
}

// New
