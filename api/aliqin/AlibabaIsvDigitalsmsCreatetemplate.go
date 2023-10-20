package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// Alibabaisvdigitalsmscreatetemplate 数字短信模板创建
// alibaba.isv.digitalsms.createtemplate
//
// 数字短信模板创建，给聚石塔，类型：2
func Alibabaisvdigitalsmscreatetemplate(clt *core.SDKClient, req *aliqin.AlibabaisvdigitalsmscreatetemplateAPIRequest, session string) (*aliqin.AlibabaisvdigitalsmscreatetemplateAPIResponse, error) {
	var resp aliqin.AlibabaisvdigitalsmscreatetemplateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
