package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscKmsAccessAPIRequest
本地生活风控数据接入 API请求
alibaba.alsc.kms.access

第三方使用本地生活数据对外提供服务，上报访问日志信息接口 */
type AlibabaAlscKmsAccessAPIRequest struct {
	model.Params
	// 接入参数
	_requestdata string
}

// New
