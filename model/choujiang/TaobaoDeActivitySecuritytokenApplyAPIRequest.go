package choujiang

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDeActivitySecuritytokenApplyAPIRequest
安全token获取 API请求
taobao.de.activity.securitytoken.apply

新增接口，这个接口是用于在手机端进行抽奖时候的验证使用 */
type TaobaoDeActivitySecuritytokenApplyAPIRequest struct {
	model.Params
	// 运营和cp约定的事件唯一标示
	_eventKey string
}

// New
