package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicecenterworkcardserviceskusuggestAPIRequest 服务商反馈需要履行的服务项 API请求
// alibaba.servicecenter.workcard.servicesku.suggest
//
// 服务商反馈需要履行的服务项
type AlibabaservicecenterworkcardserviceskusuggestAPIRequest struct {
	model.Params
	// 服务商反馈服务项
	_tpSuggestServiceSkuInfos []TpSuggestServiceSkuInfoDto
	// 真实履约服务商nick(非erp系统不要调用)
	_realTpNick string
	// 拓展信息
	_extendInfo string
	// 工单id
	_workcardId int64
}

// NewAlibabaservicecenterworkcardserviceskusuggestRequest 初始化AlibabaservicecenterworkcardserviceskusuggestAPIRequest对象
func NewAlibabaservicecenterworkcardserviceskusuggestRequest() *AlibabaservicecenterworkcardserviceskusuggestAPIRequest {
	return &AlibabaservicecenterworkcardserviceskusuggestAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaservicecenterworkcardserviceskusuggestAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.workcard.servicesku.suggest"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaservicecenterworkcardserviceskusuggestAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaservicecenterworkcardserviceskusuggestAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTpSuggestServiceSkuInfos is TpSuggestServiceSkuInfos Setter
// 服务商反馈服务项
func (r *AlibabaservicecenterworkcardserviceskusuggestAPIRequest) SetTpSuggestServiceSkuInfos(_tpSuggestServiceSkuInfos []TpSuggestServiceSkuInfoDto) error {
	r._tpSuggestServiceSkuInfos = _tpSuggestServiceSkuInfos
	r.Set("tp_suggest_service_sku_infos", _tpSuggestServiceSkuInfos)
	return nil
}

// GetTpSuggestServiceSkuInfos TpSuggestServiceSkuInfos Getter
func (r AlibabaservicecenterworkcardserviceskusuggestAPIRequest) GetTpSuggestServiceSkuInfos() []TpSuggestServiceSkuInfoDto {
	return r._tpSuggestServiceSkuInfos
}

// SetRealTpNick is RealTpNick Setter
// 真实履约服务商nick(非erp系统不要调用)
func (r *AlibabaservicecenterworkcardserviceskusuggestAPIRequest) SetRealTpNick(_realTpNick string) error {
	r._realTpNick = _realTpNick
	r.Set("real_tp_nick", _realTpNick)
	return nil
}

// GetRealTpNick RealTpNick Getter
func (r AlibabaservicecenterworkcardserviceskusuggestAPIRequest) GetRealTpNick() string {
	return r._realTpNick
}

// SetExtendInfo is ExtendInfo Setter
// 拓展信息
func (r *AlibabaservicecenterworkcardserviceskusuggestAPIRequest) SetExtendInfo(_extendInfo string) error {
	r._extendInfo = _extendInfo
	r.Set("extend_info", _extendInfo)
	return nil
}

// GetExtendInfo ExtendInfo Getter
func (r AlibabaservicecenterworkcardserviceskusuggestAPIRequest) GetExtendInfo() string {
	return r._extendInfo
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *AlibabaservicecenterworkcardserviceskusuggestAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r AlibabaservicecenterworkcardserviceskusuggestAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}
