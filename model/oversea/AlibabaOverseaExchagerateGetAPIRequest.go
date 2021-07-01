package oversea

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOverseaExchagerateGetAPIRequest
汇率信息获取 API请求
alibaba.oversea.exchagerate.get

提供外部汇率查询接口 */
type AlibabaOverseaExchagerateGetAPIRequest struct {
	model.Params
	// 业务类型
	_bizCode string
	// 原始币种
	_baseCode string
	// 目标币种
	_targetCode string
}

// New
