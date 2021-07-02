package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCbossWorkplatformWorkorderProcessNotifyAPIRequest 菜鸟工单系统的工单进度下发 API请求
// cainiao.cboss.workplatform.workorder.process.notify
//
// 菜鸟工单系统的工单进度下发（SPI）
type CainiaoCbossWorkplatformWorkorderProcessNotifyAPIRequest struct {
	model.Params
	// 服务入参
	_content *CainiaoCbossWorkplatformWorkorderProcessNotifyStruct
}

// NewCainiaoCbossWorkplatformWorkorderProcessNotifyRequest 初始化CainiaoCbossWorkplatformWorkorderProcessNotifyAPIRequest对象
func NewCainiaoCbossWorkplatformWorkorderProcessNotifyRequest() *CainiaoCbossWorkplatformWorkorderProcessNotifyAPIRequest {
	return &CainiaoCbossWorkplatformWorkorderProcessNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCbossWorkplatformWorkorderProcessNotifyAPIRequest) GetApiMethodName() string {
	return "cainiao.cboss.workplatform.workorder.process.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCbossWorkplatformWorkorderProcessNotifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetContent is Content Setter
// 服务入参
func (r *CainiaoCbossWorkplatformWorkorderProcessNotifyAPIRequest) SetContent(_content *CainiaoCbossWorkplatformWorkorderProcessNotifyStruct) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r CainiaoCbossWorkplatformWorkorderProcessNotifyAPIRequest) GetContent() *CainiaoCbossWorkplatformWorkorderProcessNotifyStruct {
	return r._content
}
