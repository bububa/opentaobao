package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkScInvitecodeGetAPIRequest
淘宝客-公用-私域用户邀请码生成 API请求
taobao.tbk.sc.invitecode.get

私域用户管理(即渠道管理或会员运营管理)功能中，通过此API可生成淘宝客自身的邀请码。 */
type TaobaoTbkScInvitecodeGetAPIRequest struct {
	model.Params
	// 渠道关系ID
	_relationId int64
	// 渠道推广的物料类型
	_relationApp string
	// 邀请码类型，1 - 渠道邀请，2 - 渠道裂变，3 -会员邀请
	_codeType int64
}

// New
