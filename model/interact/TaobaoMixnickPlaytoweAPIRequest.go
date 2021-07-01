package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMixnickPlaytoweAPIRequest
互动mixNick转微淘 API请求
taobao.mixnick.playtowe

微淘应用的混淆nick转为互动类型混淆nick */
type TaobaoMixnickPlaytoweAPIRequest struct {
	model.Params
	// 用户的混淆nick
	_mixMix string
}

// New
