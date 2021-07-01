package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeB2cMemberDiscountruleGetAPIRequest
会员抵扣规则 API请求
alibaba.nlife.b2c.member.discountrule.get

获取企业会员抵扣规则 */
type AlibabaNlifeB2cMemberDiscountruleGetAPIRequest struct {
	model.Params
	// 企业ID
	_companyId string
	// 会员在ISV处的编号
	_cardNo string
}

// New
