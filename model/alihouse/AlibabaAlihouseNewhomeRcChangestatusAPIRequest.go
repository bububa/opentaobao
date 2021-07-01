package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeRcChangestatusAPIRequest
图文草稿状态更新 API请求
alibaba.alihouse.newhome.rc.changestatus

图文草稿状态更新 */
type AlibabaAlihouseNewhomeRcChangestatusAPIRequest struct {
	model.Params
	// 外部图文id
	_outerId string
	// 0 失效 1 有效
	_status int64
}

// NewAlibabaAlihouseNewhomeRcChangestatusRequest 初始化AlibabaAlihouseNewhomeRcChangestatusAPIRequest对象
func NewAlibabaAlihouseNewhomeRcChangestatusRequest() *AlibabaAlihouseNewhomeRcChangestatusAPIRequest {
	return &AlibabaAlihouseNewhomeRcChangestatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeRcChangestatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.rc.changestatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeRcChangestatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OuterId Setter
// 外部图文id
func (r *AlibabaAlihouseNewhomeRcChangestatusAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r AlibabaAlihouseNewhomeRcChangestatusAPIRequest) GetOuterId() string {
	return r._outerId
}

// Set is Status Setter
// 0 失效 1 有效
func (r *AlibabaAlihouseNewhomeRcChangestatusAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r AlibabaAlihouseNewhomeRcChangestatusAPIRequest) GetStatus() int64 {
	return r._status
}
