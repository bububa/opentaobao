package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest 销退单上传 API请求
// taobao.rdc.aligenius.warehouse.reverse.uploading
//
// 主要用于商家上传仓库销退单信息
type TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest struct {
	model.Params
	// 参数
	_param0 *WarehouseReverseUploadingDto
}

// NewTaobaoRdcAligeniusWarehouseReverseUploadingRequest 初始化TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest对象
func NewTaobaoRdcAligeniusWarehouseReverseUploadingRequest() *TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest {
	return &TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.warehouse.reverse.uploading"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest) SetParam0(_param0 *WarehouseReverseUploadingDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest) GetParam0() *WarehouseReverseUploadingDto {
	return r._param0
}

var poolTaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRdcAligeniusWarehouseReverseUploadingRequest()
	},
}

// GetTaobaoRdcAligeniusWarehouseReverseUploadingRequest 从 sync.Pool 获取 TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest
func GetTaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest() *TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest {
	return poolTaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest.Get().(*TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest)
}

// ReleaseTaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest 将 TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest 放入 sync.Pool
func ReleaseTaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest(v *TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest) {
	v.Reset()
	poolTaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest.Put(v)
}
