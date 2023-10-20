package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcDigitalsmsCreatetemplate 数字短信模板创建
// alibaba.aliqin.fc.digitalsms.createtemplate
//
// 数字短信模板创建
func AlibabaAliqinFcDigitalsmsCreatetemplate(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest, resp *aliqin.AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
