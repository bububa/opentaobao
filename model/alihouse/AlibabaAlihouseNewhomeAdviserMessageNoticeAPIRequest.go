package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeAdviserMessageNoticeAPIRequest 催促小B发送短信 API请求
// alibaba.alihouse.newhome.adviser.message.notice
//
// 催促小B发送短信
type AlibabaAlihouseNewhomeAdviserMessageNoticeAPIRequest struct {
	model.Params
	// 置业顾问/经纪人对象
	_projectAdviserDto *ProjectAdviserDto
}

// NewAlibabaAlihouseNewhomeAdviserMessageNoticeRequest 初始化AlibabaAlihouseNewhomeAdviserMessageNoticeAPIRequest对象
func NewAlibabaAlihouseNewhomeAdviserMessageNoticeRequest() *AlibabaAlihouseNewhomeAdviserMessageNoticeAPIRequest {
	return &AlibabaAlihouseNewhomeAdviserMessageNoticeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeAdviserMessageNoticeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.adviser.message.notice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeAdviserMessageNoticeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeAdviserMessageNoticeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectAdviserDto is ProjectAdviserDto Setter
// 置业顾问/经纪人对象
func (r *AlibabaAlihouseNewhomeAdviserMessageNoticeAPIRequest) SetProjectAdviserDto(_projectAdviserDto *ProjectAdviserDto) error {
	r._projectAdviserDto = _projectAdviserDto
	r.Set("project_adviser_dto", _projectAdviserDto)
	return nil
}

// GetProjectAdviserDto ProjectAdviserDto Getter
func (r AlibabaAlihouseNewhomeAdviserMessageNoticeAPIRequest) GetProjectAdviserDto() *ProjectAdviserDto {
	return r._projectAdviserDto
}
