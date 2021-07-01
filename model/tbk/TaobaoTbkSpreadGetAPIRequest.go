package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkSpreadGetAPIRequest
淘宝客-公用-长链转短链 API请求
taobao.tbk.spread.get

输入一个原始的链接，转换得到指定的传播方式，如二维码，淘口令，短连接；
现阶段只支持短连接。 */
type TaobaoTbkSpreadGetAPIRequest struct {
	model.Params
	// 请求列表，内部包含多个url
	_requests []TbkSpreadRequest
}

// New
