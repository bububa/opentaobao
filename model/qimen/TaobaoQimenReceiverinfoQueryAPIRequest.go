package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenreceiverinfoqueryAPIRequest OAID 收件人信息解密接口 API请求
// taobao.qimen.receiverinfo.query
//
// WMS 调用该接口，通过 OAID 查询解密后的收件人信息
type TaobaoqimenreceiverinfoqueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimenreceiverinfoqueryRequest
}

// NewTaobaoqimenreceiverinfoqueryRequest 初始化TaobaoqimenreceiverinfoqueryAPIRequest对象
func NewTaobaoqimenreceiverinfoqueryRequest() *TaobaoqimenreceiverinfoqueryAPIRequest {
	return &TaobaoqimenreceiverinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenreceiverinfoqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.receiverinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenreceiverinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenreceiverinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenreceiverinfoqueryAPIRequest) SetRequest(_request *TaobaoqimenreceiverinfoqueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenreceiverinfoqueryAPIRequest) GetRequest() *TaobaoqimenreceiverinfoqueryRequest {
	return r._request
}
