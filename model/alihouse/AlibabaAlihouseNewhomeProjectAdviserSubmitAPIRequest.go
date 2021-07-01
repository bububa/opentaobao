package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest
提交楼盘顾问 API请求
alibaba.alihouse.newhome.project.adviser.submit

提交楼盘顾问 */
type AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest struct {
	model.Params
	// 顾问列表
	_advisers []ProjectAdviserDto
}

// NewAlibabaAlihouseNewhomeProjectAdviserSubmitRequest 初始化AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectAdviserSubmitRequest() *AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.adviser.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Advisers Setter
// 顾问列表
func (r *AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest) SetAdvisers(_advisers []ProjectAdviserDto) error {
	r._advisers = _advisers
	r.Set("advisers", _advisers)
	return nil
}

// Get Advisers Getter
func (r AlibabaAlihouseNewhomeProjectAdviserSubmitAPIRequest) GetAdvisers() []ProjectAdviserDto {
	return r._advisers
}
