package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeVideoChangestatusAPIRequest 视频草稿状态更新 API请求
// alibaba.alihouse.newhome.video.changestatus
//
// 视频草稿状态更新
type AlibabaAlihouseNewhomeVideoChangestatusAPIRequest struct {
	model.Params
	// 外部视频id
	_outerId string
	// 0 失效 1 有效
	_status int64
}

// NewAlibabaAlihouseNewhomeVideoChangestatusRequest 初始化AlibabaAlihouseNewhomeVideoChangestatusAPIRequest对象
func NewAlibabaAlihouseNewhomeVideoChangestatusRequest() *AlibabaAlihouseNewhomeVideoChangestatusAPIRequest {
	return &AlibabaAlihouseNewhomeVideoChangestatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.video.changestatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OuterId Setter
// 外部视频id
func (r *AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) GetOuterId() string {
	return r._outerId
}

// Set is Status Setter
// 0 失效 1 有效
func (r *AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) GetStatus() int64 {
	return r._status
}
