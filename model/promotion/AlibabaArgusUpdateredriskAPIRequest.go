package promotion

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaArgusUpdateredriskAPIRequest) GetApiMethodName() string {
	return "alibaba.argus.updateredrisk"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaArgusUpdateredriskAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RedRiskUpdateFactor Setter
// 红线价格参数
func (r *AlibabaArgusUpdateredriskAPIRequest) SetRedRiskUpdateFactor(_redRiskUpdateFactor *RedRiskUpdateFactor) error {
	r._redRiskUpdateFactor = _redRiskUpdateFactor
	r.Set("red_risk_update_factor", _redRiskUpdateFactor)
	return nil
}

// Get RedRiskUpdateFactor Getter
func (r AlibabaArgusUpdateredriskAPIRequest) GetRedRiskUpdateFactor() *RedRiskUpdateFactor {
	return r._redRiskUpdateFactor
}
