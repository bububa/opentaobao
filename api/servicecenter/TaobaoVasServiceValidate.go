package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoVasServiceValidate 增值服务订购服务验证
// taobao.vas.service.validate
//
// 增值服务订购服务验证
func TaobaoVasServiceValidate(clt *core.SDKClient, req *servicecenter.TaobaoVasServiceValidateAPIRequest, resp *servicecenter.TaobaoVasServiceValidateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
