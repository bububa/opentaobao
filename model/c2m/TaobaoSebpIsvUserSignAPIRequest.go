package c2m

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSebpIsvUserSignAPIRequest
淘小铺三方签约同步 API请求
taobao.sebp.isv.user.sign

同步淘小铺三方服务签约信息 */
type TaobaoSebpIsvUserSignAPIRequest struct {
	model.Params
	// 淘宝账号
	_userName string
	// 身份证
	_identity string
	// 到期日期
	_endTime string
	// 签约日期
	_startTime string
}

// New
