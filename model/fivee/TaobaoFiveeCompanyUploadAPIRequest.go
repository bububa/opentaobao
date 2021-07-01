package fivee

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFiveeCompanyUploadAPIRequest
上传商信息接口 API请求
taobao.fivee.company.upload

资质共享平台上传资质证照 */
type TaobaoFiveeCompanyUploadAPIRequest struct {
	model.Params
	// bu身份标识
	_paramBucode string
	// 商家证照信息
	_paramCompany *Company
}

// New
