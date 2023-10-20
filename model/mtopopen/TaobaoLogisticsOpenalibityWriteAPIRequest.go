package mtopopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsOpenalibityWriteAPIRequest 为快递公司提供的物流信息通用写入接口 API请求
// taobao.logistics.openalibity.write
//
// 为快递公司提供的物流信息通用写入接口
type TaobaoLogisticsOpenalibityWriteAPIRequest struct {
	model.Params
	// 信息写入请求对象
	_generalLogisticsDataWriteRequest *GeneralLogisticsDataWriteRequest
}

// NewTaobaoLogisticsOpenalibityWriteRequest 初始化TaobaoLogisticsOpenalibityWriteAPIRequest对象
func NewTaobaoLogisticsOpenalibityWriteRequest() *TaobaoLogisticsOpenalibityWriteAPIRequest {
	return &TaobaoLogisticsOpenalibityWriteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsOpenalibityWriteAPIRequest) Reset() {
	r._generalLogisticsDataWriteRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsOpenalibityWriteAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.openalibity.write"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsOpenalibityWriteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsOpenalibityWriteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGeneralLogisticsDataWriteRequest is GeneralLogisticsDataWriteRequest Setter
// 信息写入请求对象
func (r *TaobaoLogisticsOpenalibityWriteAPIRequest) SetGeneralLogisticsDataWriteRequest(_generalLogisticsDataWriteRequest *GeneralLogisticsDataWriteRequest) error {
	r._generalLogisticsDataWriteRequest = _generalLogisticsDataWriteRequest
	r.Set("general_logistics_data_write_request", _generalLogisticsDataWriteRequest)
	return nil
}

// GetGeneralLogisticsDataWriteRequest GeneralLogisticsDataWriteRequest Getter
func (r TaobaoLogisticsOpenalibityWriteAPIRequest) GetGeneralLogisticsDataWriteRequest() *GeneralLogisticsDataWriteRequest {
	return r._generalLogisticsDataWriteRequest
}

var poolTaobaoLogisticsOpenalibityWriteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsOpenalibityWriteRequest()
	},
}

// GetTaobaoLogisticsOpenalibityWriteRequest 从 sync.Pool 获取 TaobaoLogisticsOpenalibityWriteAPIRequest
func GetTaobaoLogisticsOpenalibityWriteAPIRequest() *TaobaoLogisticsOpenalibityWriteAPIRequest {
	return poolTaobaoLogisticsOpenalibityWriteAPIRequest.Get().(*TaobaoLogisticsOpenalibityWriteAPIRequest)
}

// ReleaseTaobaoLogisticsOpenalibityWriteAPIRequest 将 TaobaoLogisticsOpenalibityWriteAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsOpenalibityWriteAPIRequest(v *TaobaoLogisticsOpenalibityWriteAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsOpenalibityWriteAPIRequest.Put(v)
}
