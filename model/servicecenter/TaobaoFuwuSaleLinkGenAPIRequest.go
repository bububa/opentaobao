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

// NewTaobaoFuwuSaleLinkGenRequest 初始化TaobaoFuwuSaleLinkGenAPIRequest对象
func NewTaobaoFuwuSaleLinkGenRequest() *TaobaoFuwuSaleLinkGenAPIRequest {
	return &TaobaoFuwuSaleLinkGenAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFuwuSaleLinkGenAPIRequest) GetApiMethodName() string {
	return "taobao.fuwu.sale.link.gen"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFuwuSaleLinkGenAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 用户需要营销的目标人群中的用户nick
func (r *TaobaoFuwuSaleLinkGenAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoFuwuSaleLinkGenAPIRequest) GetNick() string {
	return r._nick
}

// Set is ParamStr Setter
// 从服务商后台，营销链接功能中生成的参数串直接复制使用。不要修改，否则抛错。
func (r *TaobaoFuwuSaleLinkGenAPIRequest) SetParamStr(_paramStr string) error {
	r._paramStr = _paramStr
	r.Set("param_str", _paramStr)
	return nil
}

// Get ParamStr Getter
func (r TaobaoFuwuSaleLinkGenAPIRequest) GetParamStr() string {
	return r._paramStr
}
