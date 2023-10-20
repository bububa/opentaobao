package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// Alibabaaliqinfcdigitalsmscreatetemplate 数字短信模板创建
// alibaba.aliqin.fc.digitalsms.createtemplate
//
// 数字短信模板创建
func Alibabaaliqinfcdigitalsmscreatetemplate(clt *core.SDKClient, req *aliqin.AlibabaaliqinfcdigitalsmscreatetemplateAPIRequest, session string) (*aliqin.AlibabaaliqinfcdigitalsmscreatetemplateAPIResponse, error) {
	var resp aliqin.AlibabaaliqinfcdigitalsmscreatetemplateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
