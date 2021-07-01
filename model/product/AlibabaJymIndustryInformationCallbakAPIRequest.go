package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaJymIndustryInformationCallbakAPIRequest
VMOS回调行业信息系统 API请求
alibaba.jym.industry.information.callbak

VMOS回调交易猫行业信息系统 */
type AlibabaJymIndustryInformationCallbakAPIRequest struct {
	model.Params
	// 任务ID
	_taskId string
	// 幂等ID
	_bizId string
	// 状态
	_status int64
	// 内容
	_content string
}

// New
