package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomePosApplySubmitAPIRequest pos申请接口 API请求
// alibaba.alihouse.existinghome.pos.apply.submit
//
// pos申请接口
type AlibabaAlihouseExistinghomePosApplySubmitAPIRequest struct {
	model.Params
	// 同步对象
	_sycTradePosApplyDto *SyncTradePosApplyDto
}

// NewAlibabaAlihouseExistinghomePosApplySubmitRequest 初始化AlibabaAlihouseExistinghomePosApplySubmitAPIRequest对象
func NewAlibabaAlihouseExistinghomePosApplySubmitRequest() *AlibabaAlihouseExistinghomePosApplySubmitAPIRequest {
	return &AlibabaAlihouseExistinghomePosApplySubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomePosApplySubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.pos.apply.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomePosApplySubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomePosApplySubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSycTradePosApplyDto is SycTradePosApplyDto Setter
// 同步对象
func (r *AlibabaAlihouseExistinghomePosApplySubmitAPIRequest) SetSycTradePosApplyDto(_sycTradePosApplyDto *SyncTradePosApplyDto) error {
	r._sycTradePosApplyDto = _sycTradePosApplyDto
	r.Set("syc_trade_pos_apply_dto", _sycTradePosApplyDto)
	return nil
}

// GetSycTradePosApplyDto SycTradePosApplyDto Getter
func (r AlibabaAlihouseExistinghomePosApplySubmitAPIRequest) GetSycTradePosApplyDto() *SyncTradePosApplyDto {
	return r._sycTradePosApplyDto
}
