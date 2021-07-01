package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest
提交预售证 API请求
alibaba.alihouse.newhome.project.presalepermit.submit

提交楼盘预售证信息 */
type AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIRequest struct {
	model.Params
	// 预售证对象
	_preSalePermitDto *ProjectPreSalePermitDto
}

// New
