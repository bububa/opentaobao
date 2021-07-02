package wangwang

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWangwangAbstractInitializeAPIRequest 模糊查询服务初始化 API请求
// taobao.wangwang.abstract.initialize
//
// 模糊查询服务初始化，只支持json返回
type TaobaoWangwangAbstractInitializeAPIRequest struct {
	model.Params
	// 传入参数的字符集
	_charset string
}

// NewTaobaoWangwangAbstractInitializeRequest 初始化TaobaoWangwangAbstractInitializeAPIRequest对象
func NewTaobaoWangwangAbstractInitializeRequest() *TaobaoWangwangAbstractInitializeAPIRequest {
	return &TaobaoWangwangAbstractInitializeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWangwangAbstractInitializeAPIRequest) GetApiMethodName() string {
	return "taobao.wangwang.abstract.initialize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWangwangAbstractInitializeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCharset is Charset Setter
// 传入参数的字符集
func (r *TaobaoWangwangAbstractInitializeAPIRequest) SetCharset(_charset string) error {
	r._charset = _charset
	r.Set("charset", _charset)
	return nil
}

// GetCharset Charset Getter
func (r TaobaoWangwangAbstractInitializeAPIRequest) GetCharset() string {
	return r._charset
}
