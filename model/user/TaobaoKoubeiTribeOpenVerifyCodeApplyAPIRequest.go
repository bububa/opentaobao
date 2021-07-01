package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoKoubeiTribeOpenVerifyCodeApplyAPIRequest
口碑综合体手机号获取验证码 API请求
taobao.koubei.tribe.open.verify.code.apply

口碑综合体通过手机号获取验证码对外开放接口 */
type TaobaoKoubeiTribeOpenVerifyCodeApplyAPIRequest struct {
	model.Params
	// 数据集id
	_dataSetId string
	// 手机号
	_phone string
}

// New
