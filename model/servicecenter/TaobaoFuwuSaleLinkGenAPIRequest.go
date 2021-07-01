package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFuwuSaleLinkGenAPIRequest
服务平台营销链接生成接口 API请求
taobao.fuwu.sale.link.gen

服务商通过使用该接口来产生营销链接，通过把这种链接发送给商家来做自定义人群的服务营销<br><br/>注：session是param_str这个参数串创建者生成的session，这个创建者与入参中的nick是不一致的。例如：A开发者创建了一个param_str的字符串，要为B商家生成一个营销链接，session必须是A开发者创建的session。 */
type TaobaoFuwuSaleLinkGenAPIRequest struct {
	model.Params
	// 用户需要营销的目标人群中的用户nick
	_nick string
	// 从服务商后台，营销链接功能中生成的参数串直接复制使用。不要修改，否则抛错。
	_paramStr string
}

// New
