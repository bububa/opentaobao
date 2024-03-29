package refund

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest 鉴定工单信息同步 API请求
// taobao.rdc.aligenius.identification.case.update
//
// 同步商家鉴定工单信息
type TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_param *SyncIdentifyRefundCaseDto
}

// NewTaobaoRdcAligeniusIdentificationCaseUpdateRequest 初始化TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest对象
func NewTaobaoRdcAligeniusIdentificationCaseUpdateRequest() *TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest {
	return &TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.identification.case.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求参数
func (r *TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest) SetParam(_param *SyncIdentifyRefundCaseDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest) GetParam() *SyncIdentifyRefundCaseDto {
	return r._param
}

var poolTaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRdcAligeniusIdentificationCaseUpdateRequest()
	},
}

// GetTaobaoRdcAligeniusIdentificationCaseUpdateRequest 从 sync.Pool 获取 TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest
func GetTaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest() *TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest {
	return poolTaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest.Get().(*TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest)
}

// ReleaseTaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest 将 TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest(v *TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest) {
	v.Reset()
	poolTaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest.Put(v)
}
