package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest
删除楼盘预售证 API请求
alibaba.alihouse.newhome.project.presalepermit.delete

删除楼盘预售证信息 */
type AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest struct {
	model.Params
	// 外部顾问ID
	_outerPermitId string
}

// New
