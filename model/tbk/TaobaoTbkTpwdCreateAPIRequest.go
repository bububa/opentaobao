package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkTpwdCreateAPIRequest
淘宝客-公用-淘口令生成 API请求
taobao.tbk.tpwd.create

提供淘口令生成接口。提交需要生成淘口令的内容、logo、url等参数，生成淘口令，其中关键信息为￥SADadW￥，后续可基于淘口令进行文案包装组装用于传播。 */
type TaobaoTbkTpwdCreateAPIRequest struct {
	model.Params
	// 联盟官方渠道获取的淘客推广链接，请注意，不要随意篡改官方生成的链接，否则可能无法生成淘口令
	_url string
	// 兼容旧版本api参数，无实际作用
	_text string
	// 兼容旧版本api参数，无实际作用
	_logo string
	// 兼容旧版本api参数，无实际作用
	_ext string
	// 兼容旧版本api参数，无实际作用
	_userId string
}

// New
