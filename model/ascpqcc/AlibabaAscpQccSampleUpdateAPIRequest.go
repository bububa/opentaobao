package ascpqcc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpQccSampleUpdateAPIRequest 品控中心更新样品信息 API请求
// alibaba.ascp.qcc.sample.update
//
// 品控中心更新样品信息
type AlibabaAscpQccSampleUpdateAPIRequest struct {
	model.Params
	// 更新请求参数
	_updateRequest *UpdateSampleRequest
}

// NewAlibabaAscpQccSampleUpdateRequest 初始化AlibabaAscpQccSampleUpdateAPIRequest对象
func NewAlibabaAscpQccSampleUpdateRequest() *AlibabaAscpQccSampleUpdateAPIRequest {
	return &AlibabaAscpQccSampleUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpQccSampleUpdateAPIRequest) Reset() {
	r._updateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpQccSampleUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.qcc.sample.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpQccSampleUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpQccSampleUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateRequest is UpdateRequest Setter
// 更新请求参数
func (r *AlibabaAscpQccSampleUpdateAPIRequest) SetUpdateRequest(_updateRequest *UpdateSampleRequest) error {
	r._updateRequest = _updateRequest
	r.Set("update_request", _updateRequest)
	return nil
}

// GetUpdateRequest UpdateRequest Getter
func (r AlibabaAscpQccSampleUpdateAPIRequest) GetUpdateRequest() *UpdateSampleRequest {
	return r._updateRequest
}

var poolAlibabaAscpQccSampleUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpQccSampleUpdateRequest()
	},
}

// GetAlibabaAscpQccSampleUpdateRequest 从 sync.Pool 获取 AlibabaAscpQccSampleUpdateAPIRequest
func GetAlibabaAscpQccSampleUpdateAPIRequest() *AlibabaAscpQccSampleUpdateAPIRequest {
	return poolAlibabaAscpQccSampleUpdateAPIRequest.Get().(*AlibabaAscpQccSampleUpdateAPIRequest)
}

// ReleaseAlibabaAscpQccSampleUpdateAPIRequest 将 AlibabaAscpQccSampleUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpQccSampleUpdateAPIRequest(v *AlibabaAscpQccSampleUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAscpQccSampleUpdateAPIRequest.Put(v)
}
