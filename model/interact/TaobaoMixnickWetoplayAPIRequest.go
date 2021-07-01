package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMixnickWetoplayAPIRequest
微淘混淆nick转互动混淆nick API请求
taobao.mixnick.wetoplay

微淘应用的混淆nick转为互动类型混淆nick */
type TaobaoMixnickWetoplayAPIRequest struct {
	model.Params
	// 排查问题id
	_traceId string
	// 微淘混淆nick
	_weMixnick string
}

// New
