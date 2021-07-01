package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest
提交标签库 API请求
alibaba.alihouse.newhome.base.label.submit

提交标签库 */
type AlibabaAlihouseNewhomeBaseLabelSubmitAPIRequest struct {
	model.Params
	// 标签列表
	_labels []BaseLabelDto
}

// New
