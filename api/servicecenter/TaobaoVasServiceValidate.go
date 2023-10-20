package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaovasservicevalidate 增值服务订购服务验证
// taobao.vas.service.validate
//
// 增值服务订购服务验证
func Taobaovasservicevalidate(clt *core.SDKClient, req *servicecenter.TaobaovasservicevalidateAPIRequest, session string) (*servicecenter.TaobaovasservicevalidateAPIResponse, error) {
	var resp servicecenter.TaobaovasservicevalidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
