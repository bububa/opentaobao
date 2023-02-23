package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest 服务商反馈需要履行的服务项 API请求
// alibaba.servicecenter.workcard.servicesku.suggest
//
// 服务商反馈需要履行的服务项
type AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest struct {
	model.Params
	// 服务商反馈服务项
	_tpSuggestServiceSkuInfos []TpSuggestServiceSkuInfoDto
	// 真实履约服务商nick(非erp系统不要调用)
	_realTpNick string
	// 工单id
	_workcardId int64
}

// NewAlibabaServicecenterWorkcardServiceskuSuggestRequest 初始化AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest对象
func NewAlibabaServicecenterWorkcardServiceskuSuggestRequest() *AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest {
	return &AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.workcard.servicesku.suggest"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTpSuggestServiceSkuInfos is TpSuggestServiceSkuInfos Setter
// 服务商反馈服务项
func (r *AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest) SetTpSuggestServiceSkuInfos(_tpSuggestServiceSkuInfos []TpSuggestServiceSkuInfoDto) error {
	r._tpSuggestServiceSkuInfos = _tpSuggestServiceSkuInfos
	r.Set("tp_suggest_service_sku_infos", _tpSuggestServiceSkuInfos)
	return nil
}

// GetTpSuggestServiceSkuInfos TpSuggestServiceSkuInfos Getter
func (r AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest) GetTpSuggestServiceSkuInfos() []TpSuggestServiceSkuInfoDto {
	return r._tpSuggestServiceSkuInfos
}

// SetRealTpNick is RealTpNick Setter
// 真实履约服务商nick(非erp系统不要调用)
func (r *AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest) SetRealTpNick(_realTpNick string) error {
	r._realTpNick = _realTpNick
	r.Set("real_tp_nick", _realTpNick)
	return nil
}

// GetRealTpNick RealTpNick Getter
func (r AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest) GetRealTpNick() string {
	return r._realTpNick
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}
