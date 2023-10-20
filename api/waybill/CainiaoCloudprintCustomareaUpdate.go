package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoCloudprintCustomareaUpdate 自定义区内容更新
// cainiao.cloudprint.customarea.update
//
// 自定义区内容更新
func CainiaoCloudprintCustomareaUpdate(clt *core.SDKClient, req *waybill.CainiaoCloudprintCustomareaUpdateAPIRequest, resp *waybill.CainiaoCloudprintCustomareaUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
