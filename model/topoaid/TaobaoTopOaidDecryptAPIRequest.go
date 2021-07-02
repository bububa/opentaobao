package topoaid

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopOaidDecryptAPIRequest) GetApiMethodName() string {
	return "taobao.top.oaid.decrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopOaidDecryptAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
