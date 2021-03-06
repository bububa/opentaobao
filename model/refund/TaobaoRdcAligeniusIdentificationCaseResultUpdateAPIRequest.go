package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest 鉴定工单结果同步 API请求
// taobao.rdc.aligenius.identification.case.result.update
//
// 同步鉴定工单结果信息
type TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_param *SyncIdentifyRefundCaseResultDto
}

// NewTaobaoRdcAligeniusIdentificationCaseResultUpdateRequest 初始化TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest对象
func NewTaobaoRdcAligeniusIdentificationCaseResultUpdateRequest() *TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest {
	return &TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.identification.case.result.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 请求参数
func (r *TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest) SetParam(_param *SyncIdentifyRefundCaseResultDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest) GetParam() *SyncIdentifyRefundCaseResultDto {
	return r._param
}
