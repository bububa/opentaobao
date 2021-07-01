package qt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTsSubscribeGetAPIRequest
淘宝服务订购关系查询 API请求
taobao.ts.subscribe.get

ts订购关系状态查询. 暂只支持1口价服务. */
type TaobaoTsSubscribeGetAPIRequest struct {
	model.Params
	// 服务收费项code
	_servcieItemCode string
	// 订购用户昵称
	_nick string
}

// New
