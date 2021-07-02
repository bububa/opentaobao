package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenReceiverinfoQueryAPIRequest OAID 收件人信息解密接口 API请求
// taobao.qimen.receiverinfo.query
//
// WMS 调用该接口，通过 OAID 查询解密后的收件人信息
type TaobaoQimenReceiverinfoQueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenReceiverinfoQueryRequest
}

// NewTaobaoQimenReceiverinfoQueryRequest 初始化TaobaoQimenReceiverinfoQueryAPIRequest对象
func NewTaobaoQimenReceiverinfoQueryRequest() *TaobaoQimenReceiverinfoQueryAPIRequest {
	return &TaobaoQimenReceiverinfoQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenReceiverinfoQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.receiverinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenReceiverinfoQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenReceiverinfoQueryAPIRequest) SetRequest(_request *TaobaoQimenReceiverinfoQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenReceiverinfoQueryAPIRequest) GetRequest() *TaobaoQimenReceiverinfoQueryRequest {
	return r._request
}
