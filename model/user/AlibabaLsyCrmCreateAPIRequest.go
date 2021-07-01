package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLsyCrmCreateAPIRequest
创建客资 API请求
alibaba.lsy.crm.create

欢客调用该接口进行客资创建 */
type AlibabaLsyCrmCreateAPIRequest struct {
	model.Params
	// 客资记录对象
	_nrtRecordDto *NrtRecordDto
}

// New
