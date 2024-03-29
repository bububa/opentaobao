package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest 判断支付宝用户是否绑定淘宝账号 API请求
// alibaba.aliqin.flow.alipay.isbindingtbaccount
//
// 判断支付宝用户是否绑定淘宝账号
type AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest struct {
	model.Params
	// 支付宝ID
	_alipayId string
}

// NewAlibabaAliqinFlowAlipayIsbindingtbaccountRequest 初始化AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest对象
func NewAlibabaAliqinFlowAlipayIsbindingtbaccountRequest() *AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest {
	return &AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest) Reset() {
	r._alipayId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.alipay.isbindingtbaccount"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayId is AlipayId Setter
// 支付宝ID
func (r *AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest) SetAlipayId(_alipayId string) error {
	r._alipayId = _alipayId
	r.Set("alipay_id", _alipayId)
	return nil
}

// GetAlipayId AlipayId Getter
func (r AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest) GetAlipayId() string {
	return r._alipayId
}

var poolAlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFlowAlipayIsbindingtbaccountRequest()
	},
}

// GetAlibabaAliqinFlowAlipayIsbindingtbaccountRequest 从 sync.Pool 获取 AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest
func GetAlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest() *AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest {
	return poolAlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest.Get().(*AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest)
}

// ReleaseAlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest 将 AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest(v *AlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFlowAlipayIsbindingtbaccountAPIRequest.Put(v)
}
