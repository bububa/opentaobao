package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsEfspGetuserinfoAPIRequest 获取用户详细信息 API请求
// alibaba.alisports.efsp.getuserinfo
//
// 阿里体育-体育健身-获取用户详细信息
type AlibabaAlisportsEfspGetuserinfoAPIRequest struct {
	model.Params
	// 支付宝ID
	_alipayId string
}

// NewAlibabaAlisportsEfspGetuserinfoRequest 初始化AlibabaAlisportsEfspGetuserinfoAPIRequest对象
func NewAlibabaAlisportsEfspGetuserinfoRequest() *AlibabaAlisportsEfspGetuserinfoAPIRequest {
	return &AlibabaAlisportsEfspGetuserinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsEfspGetuserinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.efsp.getuserinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsEfspGetuserinfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAlipayId is AlipayId Setter
// 支付宝ID
func (r *AlibabaAlisportsEfspGetuserinfoAPIRequest) SetAlipayId(_alipayId string) error {
	r._alipayId = _alipayId
	r.Set("alipay_id", _alipayId)
	return nil
}

// GetAlipayId AlipayId Getter
func (r AlibabaAlisportsEfspGetuserinfoAPIRequest) GetAlipayId() string {
	return r._alipayId
}
