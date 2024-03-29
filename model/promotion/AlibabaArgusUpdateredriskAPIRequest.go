package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaArgusUpdateredriskAPIRequest 更新红线价格 API请求
// alibaba.argus.updateredrisk
//
// 商品健康中心新增红线价格规则
type AlibabaArgusUpdateredriskAPIRequest struct {
	model.Params
	// 红线价格参数
	_redRiskUpdateFactor *RedRiskUpdateFactor
}

// NewAlibabaArgusUpdateredriskRequest 初始化AlibabaArgusUpdateredriskAPIRequest对象
func NewAlibabaArgusUpdateredriskRequest() *AlibabaArgusUpdateredriskAPIRequest {
	return &AlibabaArgusUpdateredriskAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaArgusUpdateredriskAPIRequest) Reset() {
	r._redRiskUpdateFactor = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaArgusUpdateredriskAPIRequest) GetApiMethodName() string {
	return "alibaba.argus.updateredrisk"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaArgusUpdateredriskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaArgusUpdateredriskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRedRiskUpdateFactor is RedRiskUpdateFactor Setter
// 红线价格参数
func (r *AlibabaArgusUpdateredriskAPIRequest) SetRedRiskUpdateFactor(_redRiskUpdateFactor *RedRiskUpdateFactor) error {
	r._redRiskUpdateFactor = _redRiskUpdateFactor
	r.Set("red_risk_update_factor", _redRiskUpdateFactor)
	return nil
}

// GetRedRiskUpdateFactor RedRiskUpdateFactor Getter
func (r AlibabaArgusUpdateredriskAPIRequest) GetRedRiskUpdateFactor() *RedRiskUpdateFactor {
	return r._redRiskUpdateFactor
}

var poolAlibabaArgusUpdateredriskAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaArgusUpdateredriskRequest()
	},
}

// GetAlibabaArgusUpdateredriskRequest 从 sync.Pool 获取 AlibabaArgusUpdateredriskAPIRequest
func GetAlibabaArgusUpdateredriskAPIRequest() *AlibabaArgusUpdateredriskAPIRequest {
	return poolAlibabaArgusUpdateredriskAPIRequest.Get().(*AlibabaArgusUpdateredriskAPIRequest)
}

// ReleaseAlibabaArgusUpdateredriskAPIRequest 将 AlibabaArgusUpdateredriskAPIRequest 放入 sync.Pool
func ReleaseAlibabaArgusUpdateredriskAPIRequest(v *AlibabaArgusUpdateredriskAPIRequest) {
	v.Reset()
	poolAlibabaArgusUpdateredriskAPIRequest.Put(v)
}
