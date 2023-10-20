package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeRcSyncAPIRequest 阿里房产图文草稿信息同步 API请求
// alibaba.alihouse.newhome.rc.sync
//
// 接收图文草稿信息
type AlibabaAlihouseNewhomeRcSyncAPIRequest struct {
	model.Params
	// 图文内容
	_rc *RichContentDraftDto
}

// NewAlibabaAlihouseNewhomeRcSyncRequest 初始化AlibabaAlihouseNewhomeRcSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeRcSyncRequest() *AlibabaAlihouseNewhomeRcSyncAPIRequest {
	return &AlibabaAlihouseNewhomeRcSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeRcSyncAPIRequest) Reset() {
	r._rc = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeRcSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.rc.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeRcSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeRcSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRc is Rc Setter
// 图文内容
func (r *AlibabaAlihouseNewhomeRcSyncAPIRequest) SetRc(_rc *RichContentDraftDto) error {
	r._rc = _rc
	r.Set("rc", _rc)
	return nil
}

// GetRc Rc Getter
func (r AlibabaAlihouseNewhomeRcSyncAPIRequest) GetRc() *RichContentDraftDto {
	return r._rc
}

var poolAlibabaAlihouseNewhomeRcSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeRcSyncRequest()
	},
}

// GetAlibabaAlihouseNewhomeRcSyncRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeRcSyncAPIRequest
func GetAlibabaAlihouseNewhomeRcSyncAPIRequest() *AlibabaAlihouseNewhomeRcSyncAPIRequest {
	return poolAlibabaAlihouseNewhomeRcSyncAPIRequest.Get().(*AlibabaAlihouseNewhomeRcSyncAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeRcSyncAPIRequest 将 AlibabaAlihouseNewhomeRcSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeRcSyncAPIRequest(v *AlibabaAlihouseNewhomeRcSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeRcSyncAPIRequest.Put(v)
}
