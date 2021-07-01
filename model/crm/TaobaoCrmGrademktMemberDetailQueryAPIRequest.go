package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmGrademktMemberDetailQueryAPIRequest
会员等级营销-等级营销活动查询 API请求
taobao.crm.grademkt.member.detail.query

商家通过该接口查询等级营销活动 */
type TaobaoCrmGrademktMemberDetailQueryAPIRequest struct {
	model.Params
	// 扩展字段
	_feather string
	// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
	_parameter string
}

// New
