package mei

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCrmMemberPointChangeAPIRequest
会员积分变更 API请求
tmall.crm.member.point.change

品牌CRM项目中，会员积分变更接口。 */
type TmallCrmMemberPointChangeAPIRequest struct {
	model.Params
	// 更改积分数值
	_point int64
	// minus:扣减;add:累加
	_type string
	// 业务代码
	_bizCode string
	// 业务描述
	_bizDetail string
	// 用户昵称
	_userNick string
}

// New
