package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeScanHisAPIRequest 企业查询扫码历史 API请求
// alibaba.alihealth.drugcode.scan.his
//
// 企业查询扫码历史
type AlibabaAlihealthDrugcodeScanHisAPIRequest struct {
	model.Params
	// 追溯码
	_code string
}

// NewAlibabaAlihealthDrugcodeScanHisRequest 初始化AlibabaAlihealthDrugcodeScanHisAPIRequest对象
func NewAlibabaAlihealthDrugcodeScanHisRequest() *AlibabaAlihealthDrugcodeScanHisAPIRequest {
	return &AlibabaAlihealthDrugcodeScanHisAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugcodeScanHisAPIRequest) Reset() {
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeScanHisAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.scan.his"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeScanHisAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugcodeScanHisAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugcodeScanHisAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugcodeScanHisAPIRequest) GetCode() string {
	return r._code
}

var poolAlibabaAlihealthDrugcodeScanHisAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugcodeScanHisRequest()
	},
}

// GetAlibabaAlihealthDrugcodeScanHisRequest 从 sync.Pool 获取 AlibabaAlihealthDrugcodeScanHisAPIRequest
func GetAlibabaAlihealthDrugcodeScanHisAPIRequest() *AlibabaAlihealthDrugcodeScanHisAPIRequest {
	return poolAlibabaAlihealthDrugcodeScanHisAPIRequest.Get().(*AlibabaAlihealthDrugcodeScanHisAPIRequest)
}

// ReleaseAlibabaAlihealthDrugcodeScanHisAPIRequest 将 AlibabaAlihealthDrugcodeScanHisAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeScanHisAPIRequest(v *AlibabaAlihealthDrugcodeScanHisAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeScanHisAPIRequest.Put(v)
}
