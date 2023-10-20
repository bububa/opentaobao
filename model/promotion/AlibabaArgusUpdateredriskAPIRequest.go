package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaargusupdateredriskAPIRequest 更新红线价格 API请求
// alibaba.argus.updateredrisk
//
// 商品健康中心新增红线价格规则
type AlibabaargusupdateredriskAPIRequest struct {
	model.Params
	// 红线价格参数
	_redRiskUpdateFactor *RedRiskUpdateFactor
}

// NewAlibabaargusupdateredriskRequest 初始化AlibabaargusupdateredriskAPIRequest对象
func NewAlibabaargusupdateredriskRequest() *AlibabaargusupdateredriskAPIRequest {
	return &AlibabaargusupdateredriskAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaargusupdateredriskAPIRequest) GetApiMethodName() string {
	return "alibaba.argus.updateredrisk"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaargusupdateredriskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaargusupdateredriskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRedRiskUpdateFactor is RedRiskUpdateFactor Setter
// 红线价格参数
func (r *AlibabaargusupdateredriskAPIRequest) SetRedRiskUpdateFactor(_redRiskUpdateFactor *RedRiskUpdateFactor) error {
	r._redRiskUpdateFactor = _redRiskUpdateFactor
	r.Set("red_risk_update_factor", _redRiskUpdateFactor)
	return nil
}

// GetRedRiskUpdateFactor RedRiskUpdateFactor Getter
func (r AlibabaargusupdateredriskAPIRequest) GetRedRiskUpdateFactor() *RedRiskUpdateFactor {
	return r._redRiskUpdateFactor
}
