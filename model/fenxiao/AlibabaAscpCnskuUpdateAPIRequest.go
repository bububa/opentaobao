package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpCnskuUpdateAPIRequest
供应链中台货品修改接口 API请求
alibaba.ascp.cnsku.update

供应链中台货品修改接口 */
type AlibabaAscpCnskuUpdateAPIRequest struct {
	model.Params
	// 待新增的货品
	_cnsku *CnskuDto
	// 修改选项
	_option *UpdateCnskuOption
}

// New
