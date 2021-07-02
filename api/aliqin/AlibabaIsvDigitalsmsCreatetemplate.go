package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaIsvDigitalsmsCreatetemplate 数字短信模板创建
// alibaba.isv.digitalsms.createtemplate
//
// 数字短信模板创建，给聚石塔，类型：2
func AlibabaIsvDigitalsmsCreatetemplate(clt *core.SDKClient, req *aliqin.AlibabaIsvDigitalsmsCreatetemplateAPIRequest, session string) (*aliqin.AlibabaIsvDigitalsmsCreatetemplateAPIResponse, error) {
	var resp aliqin.AlibabaIsvDigitalsmsCreatetemplateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
