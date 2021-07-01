package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmGradeGetAPIRequest
卖家查询等级规则 API请求
taobao.crm.grade.get

卖家查询等级规则，包括店铺客户、普通会员、高级会员、VIP会员、至尊VIP会员四个等级的信息 */
type TaobaoCrmGradeGetAPIRequest struct {
	model.Params
}

// New
