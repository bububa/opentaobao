package tbtrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopOaidDecryptAPIRequest OAID解密 API请求
// taobao.top.oaid.decrypt
//
// 解码OAID(Open Addressee ID)，返回收件人信息。
type TaobaoTopOaidDecryptAPIRequest struct {
	model.Params
	// 解密请求列表，最多支持20个。
	_queryList []ReceiverQuery
}

// NewTaobaoTopOaidDecryptRequest 初始化TaobaoTopOaidDecryptAPIRequest对象
func NewTaobaoTopOaidDecryptRequest() *TaobaoTopOaidDecryptAPIRequest {
	return &TaobaoTopOaidDecryptAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTopOaidDecryptAPIRequest) Reset() {
	r._queryList = r._queryList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopOaidDecryptAPIRequest) GetApiMethodName() string {
	return "taobao.top.oaid.decrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopOaidDecryptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopOaidDecryptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryList is QueryList Setter
// 解密请求列表，最多支持20个。
func (r *TaobaoTopOaidDecryptAPIRequest) SetQueryList(_queryList []ReceiverQuery) error {
	r._queryList = _queryList
	r.Set("query_list", _queryList)
	return nil
}

// GetQueryList QueryList Getter
func (r TaobaoTopOaidDecryptAPIRequest) GetQueryList() []ReceiverQuery {
	return r._queryList
}

var poolTaobaoTopOaidDecryptAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTopOaidDecryptRequest()
	},
}

// GetTaobaoTopOaidDecryptRequest 从 sync.Pool 获取 TaobaoTopOaidDecryptAPIRequest
func GetTaobaoTopOaidDecryptAPIRequest() *TaobaoTopOaidDecryptAPIRequest {
	return poolTaobaoTopOaidDecryptAPIRequest.Get().(*TaobaoTopOaidDecryptAPIRequest)
}

// ReleaseTaobaoTopOaidDecryptAPIRequest 将 TaobaoTopOaidDecryptAPIRequest 放入 sync.Pool
func ReleaseTaobaoTopOaidDecryptAPIRequest(v *TaobaoTopOaidDecryptAPIRequest) {
	v.Reset()
	poolTaobaoTopOaidDecryptAPIRequest.Put(v)
}
