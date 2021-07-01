package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthPregnancyBaseinfoGetAPIRequest
拉取备孕初始化信息 API请求
alibaba.alihealth.pregnancy.baseinfo.get

拉取备孕初始化信息 */
type AlibabaAlihealthPregnancyBaseinfoGetAPIRequest struct {
	model.Params
	// 用户id
	_userId int64
}

// New
