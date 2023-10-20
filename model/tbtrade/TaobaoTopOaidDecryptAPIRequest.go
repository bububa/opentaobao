package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopoaiddecryptAPIRequest OAID解密 API请求
// taobao.top.oaid.decrypt
//
// 解码OAID(Open Addressee ID)，返回收件人信息。
type TaobaotopoaiddecryptAPIRequest struct {
	model.Params
	// 解密请求列表，最多支持20个。
	_queryList []ReceiverQuery
}

// NewTaobaotopoaiddecryptRequest 初始化TaobaotopoaiddecryptAPIRequest对象
func NewTaobaotopoaiddecryptRequest() *TaobaotopoaiddecryptAPIRequest {
	return &TaobaotopoaiddecryptAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotopoaiddecryptAPIRequest) GetApiMethodName() string {
	return "taobao.top.oaid.decrypt"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotopoaiddecryptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotopoaiddecryptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryList is QueryList Setter
// 解密请求列表，最多支持20个。
func (r *TaobaotopoaiddecryptAPIRequest) SetQueryList(_queryList []ReceiverQuery) error {
	r._queryList = _queryList
	r.Set("query_list", _queryList)
	return nil
}

// GetQueryList QueryList Getter
func (r TaobaotopoaiddecryptAPIRequest) GetQueryList() []ReceiverQuery {
	return r._queryList
}
