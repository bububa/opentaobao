package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthbcitemperiodsyncAPIRequest 代销品效期同步 API请求
// alibaba.alihealth.bc.item.period.sync
//
// 代销品效期同步
type AlibabaalihealthbcitemperiodsyncAPIRequest struct {
	model.Params
	// 请求体
	_validityPeriodSyncReqDto *ValidityPeriodSyncReqDto
}

// NewAlibabaalihealthbcitemperiodsyncRequest 初始化AlibabaalihealthbcitemperiodsyncAPIRequest对象
func NewAlibabaalihealthbcitemperiodsyncRequest() *AlibabaalihealthbcitemperiodsyncAPIRequest {
	return &AlibabaalihealthbcitemperiodsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthbcitemperiodsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.bc.item.period.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthbcitemperiodsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthbcitemperiodsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetValidityPeriodSyncReqDto is ValidityPeriodSyncReqDto Setter
// 请求体
func (r *AlibabaalihealthbcitemperiodsyncAPIRequest) SetValidityPeriodSyncReqDto(_validityPeriodSyncReqDto *ValidityPeriodSyncReqDto) error {
	r._validityPeriodSyncReqDto = _validityPeriodSyncReqDto
	r.Set("validity_period_sync_req_dto", _validityPeriodSyncReqDto)
	return nil
}

// GetValidityPeriodSyncReqDto ValidityPeriodSyncReqDto Getter
func (r AlibabaalihealthbcitemperiodsyncAPIRequest) GetValidityPeriodSyncReqDto() *ValidityPeriodSyncReqDto {
	return r._validityPeriodSyncReqDto
}
