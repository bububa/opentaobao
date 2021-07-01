package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmGrademktMemberQueryAPIRequest
会员等级营销-会员关系查询 API请求
taobao.crm.grademkt.member.query

商家通过该接口查询线上店铺会员。 */
type TaobaoCrmGrademktMemberQueryAPIRequest struct {
	model.Params
	// 会员属性-json format生成方法见http://open.taobao.com/doc/detail.htm?id=101281
	_parameter string
	// 系统属性，json格式
	_feather string
	// 会员nick
	_buyerNick string
}

// New
