package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomeposapplysubmitAPIRequest pos申请接口 API请求
// alibaba.alihouse.existinghome.pos.apply.submit
//
// pos申请接口
type AlibabaalihouseexistinghomeposapplysubmitAPIRequest struct {
	model.Params
	// 同步对象
	_sycTradePosApplyDto *SyncTradePosApplyDto
}

// NewAlibabaalihouseexistinghomeposapplysubmitRequest 初始化AlibabaalihouseexistinghomeposapplysubmitAPIRequest对象
func NewAlibabaalihouseexistinghomeposapplysubmitRequest() *AlibabaalihouseexistinghomeposapplysubmitAPIRequest {
	return &AlibabaalihouseexistinghomeposapplysubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomeposapplysubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.pos.apply.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomeposapplysubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomeposapplysubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSycTradePosApplyDto is SycTradePosApplyDto Setter
// 同步对象
func (r *AlibabaalihouseexistinghomeposapplysubmitAPIRequest) SetSycTradePosApplyDto(_sycTradePosApplyDto *SyncTradePosApplyDto) error {
	r._sycTradePosApplyDto = _sycTradePosApplyDto
	r.Set("syc_trade_pos_apply_dto", _sycTradePosApplyDto)
	return nil
}

// GetSycTradePosApplyDto SycTradePosApplyDto Getter
func (r AlibabaalihouseexistinghomeposapplysubmitAPIRequest) GetSycTradePosApplyDto() *SyncTradePosApplyDto {
	return r._sycTradePosApplyDto
}
