package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoCloudprintMystdtemplatesGet 获取用户使用的菜鸟电子面单模板信息
// cainiao.cloudprint.mystdtemplates.get
//
// 获取用户使用的菜鸟电子面单
func CainiaoCloudprintMystdtemplatesGet(clt *core.SDKClient, req *waybill.CainiaoCloudprintMystdtemplatesGetAPIRequest, resp *waybill.CainiaoCloudprintMystdtemplatesGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
