package tmallsc

import (
	"net/url"
	"sync"

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
	// 拓展信息
	_extendInfo string
	// 工单id
	_workcardId int64
}

// NewAlibabaServicecenterWorkcardServiceskuSuggestRequest 初始化AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest对象
func NewAlibabaServicecenterWorkcardServiceskuSuggestRequest() *AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest {
	return &AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest) Reset() {
	r._tpSuggestServiceSkuInfos = r._tpSuggestServiceSkuInfos[:0]
	r._realTpNick = ""
	r._extendInfo = ""
	r._workcardId = 0
	r.Params.ToZero()
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

// SetExtendInfo is ExtendInfo Setter
// 拓展信息
func (r *AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest) SetExtendInfo(_extendInfo string) error {
	r._extendInfo = _extendInfo
	r.Set("extend_info", _extendInfo)
	return nil
}

// GetExtendInfo ExtendInfo Getter
func (r AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest) GetExtendInfo() string {
	return r._extendInfo
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

var poolAlibabaServicecenterWorkcardServiceskuSuggestAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaServicecenterWorkcardServiceskuSuggestRequest()
	},
}

// GetAlibabaServicecenterWorkcardServiceskuSuggestRequest 从 sync.Pool 获取 AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest
func GetAlibabaServicecenterWorkcardServiceskuSuggestAPIRequest() *AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest {
	return poolAlibabaServicecenterWorkcardServiceskuSuggestAPIRequest.Get().(*AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest)
}

// ReleaseAlibabaServicecenterWorkcardServiceskuSuggestAPIRequest 将 AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest 放入 sync.Pool
func ReleaseAlibabaServicecenterWorkcardServiceskuSuggestAPIRequest(v *AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest) {
	v.Reset()
	poolAlibabaServicecenterWorkcardServiceskuSuggestAPIRequest.Put(v)
}
