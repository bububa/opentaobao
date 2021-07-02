package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeRcSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.rc.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeRcSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
