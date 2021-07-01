package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFiveeCompanyGetAPIRequest
查询商信息 API请求
taobao.fivee.company.get

资质共享平台查询商信息 */
type TaobaoFiveeCompanyGetAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBucode string
	// 统一社会信息用代码
	_paramUniqueCode string
}

// New
