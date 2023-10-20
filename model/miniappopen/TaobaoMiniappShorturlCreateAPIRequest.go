package miniappopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappShorturlCreateAPIRequest 生成淘宝小程序短链接 API请求
// taobao.miniapp.shorturl.create
//
// 提供淘宝小程序短链接生成的能力，只允许对淘宝小程序对应的域名：https://m.duanqu.com/ 生成对应的短链接，其他域名无效
// 【特别注意：短链接有效期为30天，超过时效短链接将无法正常跳转到原始链接地址，请勿将短链接投放或装修到长期存在的入口】
type TaobaoMiniappShorturlCreateAPIRequest struct {
	model.Params
	// 小程序链接地址。说明：链接地址，只允许https协议，域名只支持m.duanqu.com，链接必须包含_ariver_appid参数，链接不能够包含spm、short_name、app、tb_url_time_stamp这些系统保留参数
	_miniappUrl string
}

// NewTaobaoMiniappShorturlCreateRequest 初始化TaobaoMiniappShorturlCreateAPIRequest对象
func NewTaobaoMiniappShorturlCreateRequest() *TaobaoMiniappShorturlCreateAPIRequest {
	return &TaobaoMiniappShorturlCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappShorturlCreateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.shorturl.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappShorturlCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappShorturlCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMiniappUrl is MiniappUrl Setter
// 小程序链接地址。说明：链接地址，只允许https协议，域名只支持m.duanqu.com，链接必须包含_ariver_appid参数，链接不能够包含spm、short_name、app、tb_url_time_stamp这些系统保留参数
func (r *TaobaoMiniappShorturlCreateAPIRequest) SetMiniappUrl(_miniappUrl string) error {
	r._miniappUrl = _miniappUrl
	r.Set("miniapp_url", _miniappUrl)
	return nil
}

// GetMiniappUrl MiniappUrl Getter
func (r TaobaoMiniappShorturlCreateAPIRequest) GetMiniappUrl() string {
	return r._miniappUrl
}
