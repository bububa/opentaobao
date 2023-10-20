package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintClientinfoPutAPIRequest 云打印客户端监控信息收集 API请求
// cainiao.cloudprint.clientinfo.put
//
// 云打印客户端监控信息收集
type CainiaoCloudprintClientinfoPutAPIRequest struct {
	model.Params
	// 客户端上传json数据
	_jsonData string
}

// NewCainiaoCloudprintClientinfoPutRequest 初始化CainiaoCloudprintClientinfoPutAPIRequest对象
func NewCainiaoCloudprintClientinfoPutRequest() *CainiaoCloudprintClientinfoPutAPIRequest {
	return &CainiaoCloudprintClientinfoPutAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoCloudprintClientinfoPutAPIRequest) Reset() {
	r._jsonData = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintClientinfoPutAPIRequest) GetApiMethodName() string {
	return "cainiao.cloudprint.clientinfo.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintClientinfoPutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCloudprintClientinfoPutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetJsonData is JsonData Setter
// 客户端上传json数据
func (r *CainiaoCloudprintClientinfoPutAPIRequest) SetJsonData(_jsonData string) error {
	r._jsonData = _jsonData
	r.Set("json_data", _jsonData)
	return nil
}

// GetJsonData JsonData Getter
func (r CainiaoCloudprintClientinfoPutAPIRequest) GetJsonData() string {
	return r._jsonData
}

var poolCainiaoCloudprintClientinfoPutAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoCloudprintClientinfoPutRequest()
	},
}

// GetCainiaoCloudprintClientinfoPutRequest 从 sync.Pool 获取 CainiaoCloudprintClientinfoPutAPIRequest
func GetCainiaoCloudprintClientinfoPutAPIRequest() *CainiaoCloudprintClientinfoPutAPIRequest {
	return poolCainiaoCloudprintClientinfoPutAPIRequest.Get().(*CainiaoCloudprintClientinfoPutAPIRequest)
}

// ReleaseCainiaoCloudprintClientinfoPutAPIRequest 将 CainiaoCloudprintClientinfoPutAPIRequest 放入 sync.Pool
func ReleaseCainiaoCloudprintClientinfoPutAPIRequest(v *CainiaoCloudprintClientinfoPutAPIRequest) {
	v.Reset()
	poolCainiaoCloudprintClientinfoPutAPIRequest.Put(v)
}
