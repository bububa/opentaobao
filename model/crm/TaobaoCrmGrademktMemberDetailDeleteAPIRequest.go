package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmGrademktMemberDetailDeleteAPIRequest
会员等级营销-删除商品等级营销明细 API请求
taobao.crm.grademkt.member.detail.delete

删除商品等级营销明细 */
type TaobaoCrmGrademktMemberDetailDeleteAPIRequest struct {
	model.Params
	// 扩展字段
	_feather string
	// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
	_parameter string
}

// New
