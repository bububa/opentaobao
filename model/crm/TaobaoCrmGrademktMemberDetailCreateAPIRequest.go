package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmGrademktMemberDetailCreateAPIRequest
会员等级营销-创建商品等级营销明细 API请求
taobao.crm.grademkt.member.detail.create

创建商品等级营销明细 */
type TaobaoCrmGrademktMemberDetailCreateAPIRequest struct {
	model.Params
	// 扩展字段
	_feather string
	// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
	_parameter string
}

// New
