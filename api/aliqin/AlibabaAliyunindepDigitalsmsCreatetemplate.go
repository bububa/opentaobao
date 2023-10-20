package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// Alibabaaliyunindepdigitalsmscreatetemplate 数字短信模板创建
// alibaba.aliyunindep.digitalsms.createtemplate
//
// 数字短信模板创建，给阿里云一方产品使用，类型：9
func Alibabaaliyunindepdigitalsmscreatetemplate(clt *core.SDKClient, req *aliqin.AlibabaaliyunindepdigitalsmscreatetemplateAPIRequest, session string) (*aliqin.AlibabaaliyunindepdigitalsmscreatetemplateAPIResponse, error) {
	var resp aliqin.AlibabaaliyunindepdigitalsmscreatetemplateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
