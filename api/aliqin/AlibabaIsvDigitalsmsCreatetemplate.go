package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaIsvDigitalsmsCreatetemplate 数字短信模板创建
// alibaba.isv.digitalsms.createtemplate
//
// 数字短信模板创建，给聚石塔，类型：2
func AlibabaIsvDigitalsmsCreatetemplate(clt *core.SDKClient, req *aliqin.AlibabaIsvDigitalsmsCreatetemplateAPIRequest, resp *aliqin.AlibabaIsvDigitalsmsCreatetemplateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
