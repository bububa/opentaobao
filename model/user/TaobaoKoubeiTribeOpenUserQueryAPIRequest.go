package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoKoubeiTribeOpenUserQueryAPIRequest
获取用户openId API请求
taobao.koubei.tribe.open.user.query

口碑综合体通过手机号码获取加密后的用户openId */
type TaobaoKoubeiTribeOpenUserQueryAPIRequest struct {
	model.Params
	// 验证码
	_verifyCode string
	// 手机号
	_phone string
	// 数据集id
	_dataSetId string
}

// New
