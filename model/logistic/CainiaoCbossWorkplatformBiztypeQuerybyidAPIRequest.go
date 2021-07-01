package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCbossWorkplatformBiztypeQuerybyidAPIRequest
菜鸟工单平台根据业务类型id查询业务类型详细信息 API请求
cainiao.cboss.workplatform.biztype.querybyid

菜鸟工单平台根据业务类型id查询业务类型详细信息。 目前调用者ISV */
type CainiaoCbossWorkplatformBiztypeQuerybyidAPIRequest struct {
	model.Params
	// 业务类型id
	_bizTypeId string
}

// New
