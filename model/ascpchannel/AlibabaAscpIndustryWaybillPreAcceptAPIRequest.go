package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryWaybillPreAcceptAPIRequest 商家ERP预推单 API请求
// alibaba.ascp.industry.waybill.pre.accept
//
// 商家ERP预推单
type AlibabaAscpIndustryWaybillPreAcceptAPIRequest struct {
	model.Params
	// 开单信息
	_waybillGenRequest *WaybillGenRequest
}

// NewAlibabaAscpIndustryWaybillPreAcceptRequest 初始化AlibabaAscpIndustryWaybillPreAcceptAPIRequest对象
func NewAlibabaAscpIndustryWaybillPreAcceptRequest() *AlibabaAscpIndustryWaybillPreAcceptAPIRequest {
	return &AlibabaAscpIndustryWaybillPreAcceptAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpIndustryWaybillPreAcceptAPIRequest) Reset() {
	r._waybillGenRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryWaybillPreAcceptAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.waybill.pre.accept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryWaybillPreAcceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpIndustryWaybillPreAcceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillGenRequest is WaybillGenRequest Setter
// 开单信息
func (r *AlibabaAscpIndustryWaybillPreAcceptAPIRequest) SetWaybillGenRequest(_waybillGenRequest *WaybillGenRequest) error {
	r._waybillGenRequest = _waybillGenRequest
	r.Set("waybill_gen_request", _waybillGenRequest)
	return nil
}

// GetWaybillGenRequest WaybillGenRequest Getter
func (r AlibabaAscpIndustryWaybillPreAcceptAPIRequest) GetWaybillGenRequest() *WaybillGenRequest {
	return r._waybillGenRequest
}

var poolAlibabaAscpIndustryWaybillPreAcceptAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpIndustryWaybillPreAcceptRequest()
	},
}

// GetAlibabaAscpIndustryWaybillPreAcceptRequest 从 sync.Pool 获取 AlibabaAscpIndustryWaybillPreAcceptAPIRequest
func GetAlibabaAscpIndustryWaybillPreAcceptAPIRequest() *AlibabaAscpIndustryWaybillPreAcceptAPIRequest {
	return poolAlibabaAscpIndustryWaybillPreAcceptAPIRequest.Get().(*AlibabaAscpIndustryWaybillPreAcceptAPIRequest)
}

// ReleaseAlibabaAscpIndustryWaybillPreAcceptAPIRequest 将 AlibabaAscpIndustryWaybillPreAcceptAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpIndustryWaybillPreAcceptAPIRequest(v *AlibabaAscpIndustryWaybillPreAcceptAPIRequest) {
	v.Reset()
	poolAlibabaAscpIndustryWaybillPreAcceptAPIRequest.Put(v)
}
