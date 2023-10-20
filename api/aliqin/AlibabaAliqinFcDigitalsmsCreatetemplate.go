package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcDigitalsmsCreatetemplate 数字短信模板创建
// alibaba.aliqin.fc.digitalsms.createtemplate
//
// 数字短信模板创建
func AlibabaAliqinFcDigitalsmsCreatetemplate(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcDigitalsmsCreatetemplateAPIRequest, session string) (*aliqin.AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse, error) {
	var resp aliqin.AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
