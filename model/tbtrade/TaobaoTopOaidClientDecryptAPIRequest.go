package tbtrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopOaidClientDecryptAPIRequest 端侧OAID解密 API请求
// taobao.top.oaid.client.decrypt
//
// 解码OAID(Open Addressee ID)，返回收件人信息。该接口用于客户端直接查看订单隐私数据，解密数据不经过ISV服务器，且包含风控等安全检测。
type TaobaoTopOaidClientDecryptAPIRequest struct {
	model.Params
	// 解密请求列表，长度不要超过1（只能单笔解密，不支持批量解密）。
	_queryList []ReceiverQuery
	// 安全令牌
	_secOnceToken string
}

// NewTaobaoTopOaidClientDecryptRequest 初始化TaobaoTopOaidClientDecryptAPIRequest对象
func NewTaobaoTopOaidClientDecryptRequest() *TaobaoTopOaidClientDecryptAPIRequest {
	return &TaobaoTopOaidClientDecryptAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTopOaidClientDecryptAPIRequest) Reset() {
	r._queryList = r._queryList[:0]
	r._secOnceToken = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopOaidClientDecryptAPIRequest) GetApiMethodName() string {
	return "taobao.top.oaid.client.decrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopOaidClientDecryptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopOaidClientDecryptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryList is QueryList Setter
// 解密请求列表，长度不要超过1（只能单笔解密，不支持批量解密）。
func (r *TaobaoTopOaidClientDecryptAPIRequest) SetQueryList(_queryList []ReceiverQuery) error {
	r._queryList = _queryList
	r.Set("query_list", _queryList)
	return nil
}

// GetQueryList QueryList Getter
func (r TaobaoTopOaidClientDecryptAPIRequest) GetQueryList() []ReceiverQuery {
	return r._queryList
}

// SetSecOnceToken is SecOnceToken Setter
// 安全令牌
func (r *TaobaoTopOaidClientDecryptAPIRequest) SetSecOnceToken(_secOnceToken string) error {
	r._secOnceToken = _secOnceToken
	r.Set("sec_once_token", _secOnceToken)
	return nil
}

// GetSecOnceToken SecOnceToken Getter
func (r TaobaoTopOaidClientDecryptAPIRequest) GetSecOnceToken() string {
	return r._secOnceToken
}

var poolTaobaoTopOaidClientDecryptAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTopOaidClientDecryptRequest()
	},
}

// GetTaobaoTopOaidClientDecryptRequest 从 sync.Pool 获取 TaobaoTopOaidClientDecryptAPIRequest
func GetTaobaoTopOaidClientDecryptAPIRequest() *TaobaoTopOaidClientDecryptAPIRequest {
	return poolTaobaoTopOaidClientDecryptAPIRequest.Get().(*TaobaoTopOaidClientDecryptAPIRequest)
}

// ReleaseTaobaoTopOaidClientDecryptAPIRequest 将 TaobaoTopOaidClientDecryptAPIRequest 放入 sync.Pool
func ReleaseTaobaoTopOaidClientDecryptAPIRequest(v *TaobaoTopOaidClientDecryptAPIRequest) {
	v.Reset()
	poolTaobaoTopOaidClientDecryptAPIRequest.Put(v)
}
