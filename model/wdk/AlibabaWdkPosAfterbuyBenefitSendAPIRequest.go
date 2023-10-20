package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPosAfterbuyBenefitSendAPIRequest 生态pos购后发放权益 API请求
// alibaba.wdk.pos.afterbuy.benefit.send
//
// 生态pos购后发放权益接口开放
type AlibabaWdkPosAfterbuyBenefitSendAPIRequest struct {
	model.Params
	// 入参
	_sendBenefitParam *IsvSendBenefitParam
}

// NewAlibabaWdkPosAfterbuyBenefitSendRequest 初始化AlibabaWdkPosAfterbuyBenefitSendAPIRequest对象
func NewAlibabaWdkPosAfterbuyBenefitSendRequest() *AlibabaWdkPosAfterbuyBenefitSendAPIRequest {
	return &AlibabaWdkPosAfterbuyBenefitSendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkPosAfterbuyBenefitSendAPIRequest) Reset() {
	r._sendBenefitParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosAfterbuyBenefitSendAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.afterbuy.benefit.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosAfterbuyBenefitSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkPosAfterbuyBenefitSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSendBenefitParam is SendBenefitParam Setter
// 入参
func (r *AlibabaWdkPosAfterbuyBenefitSendAPIRequest) SetSendBenefitParam(_sendBenefitParam *IsvSendBenefitParam) error {
	r._sendBenefitParam = _sendBenefitParam
	r.Set("send_benefit_param", _sendBenefitParam)
	return nil
}

// GetSendBenefitParam SendBenefitParam Getter
func (r AlibabaWdkPosAfterbuyBenefitSendAPIRequest) GetSendBenefitParam() *IsvSendBenefitParam {
	return r._sendBenefitParam
}

var poolAlibabaWdkPosAfterbuyBenefitSendAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkPosAfterbuyBenefitSendRequest()
	},
}

// GetAlibabaWdkPosAfterbuyBenefitSendRequest 从 sync.Pool 获取 AlibabaWdkPosAfterbuyBenefitSendAPIRequest
func GetAlibabaWdkPosAfterbuyBenefitSendAPIRequest() *AlibabaWdkPosAfterbuyBenefitSendAPIRequest {
	return poolAlibabaWdkPosAfterbuyBenefitSendAPIRequest.Get().(*AlibabaWdkPosAfterbuyBenefitSendAPIRequest)
}

// ReleaseAlibabaWdkPosAfterbuyBenefitSendAPIRequest 将 AlibabaWdkPosAfterbuyBenefitSendAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkPosAfterbuyBenefitSendAPIRequest(v *AlibabaWdkPosAfterbuyBenefitSendAPIRequest) {
	v.Reset()
	poolAlibabaWdkPosAfterbuyBenefitSendAPIRequest.Put(v)
}
