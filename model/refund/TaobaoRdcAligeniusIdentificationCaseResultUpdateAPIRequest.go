package refund

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.identification.case.result.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRdcAligeniusIdentificationCaseResultUpdateRequest()
	},
}

// GetTaobaoRdcAligeniusIdentificationCaseResultUpdateRequest 从 sync.Pool 获取 TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest
func GetTaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest() *TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest {
	return poolTaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest.Get().(*TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest)
}

// ReleaseTaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest 将 TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest(v *TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest) {
	v.Reset()
	poolTaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest.Put(v)
}
