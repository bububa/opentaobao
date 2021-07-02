package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectAdviserDeleteAPIRequest 删除楼盘顾问 API请求
// alibaba.alihouse.newhome.project.adviser.delete
//
// 删除楼盘顾问
type AlibabaAlihouseNewhomeProjectAdviserDeleteAPIRequest struct {
	model.Params
	// 外部顾问ID
	_outerConsultantId string
}

// NewAlibabaAlihouseNewhomeProjectAdviserDeleteRequest 初始化AlibabaAlihouseNewhomeProjectAdviserDeleteAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectAdviserDeleteRequest() *AlibabaAlihouseNewhomeProjectAdviserDeleteAPIRequest {
	return &AlibabaAlihouseNewhomeProjectAdviserDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectAdviserDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.adviser.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectAdviserDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OuterConsultantId Setter
// 外部顾问ID
func (r *AlibabaAlihouseNewhomeProjectAdviserDeleteAPIRequest) SetOuterConsultantId(_outerConsultantId string) error {
	r._outerConsultantId = _outerConsultantId
	r.Set("outer_consultant_id", _outerConsultantId)
	return nil
}

// Get OuterConsultantId Getter
func (r AlibabaAlihouseNewhomeProjectAdviserDeleteAPIRequest) GetOuterConsultantId() string {
	return r._outerConsultantId
}
