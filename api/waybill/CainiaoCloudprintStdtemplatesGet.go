package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoCloudprintStdtemplatesGet 获取所有的菜鸟标准电子面单模板
// cainiao.cloudprint.stdtemplates.get
//
// 获取菜鸟标准电子面单模板
func CainiaoCloudprintStdtemplatesGet(clt *core.SDKClient, req *waybill.CainiaoCloudprintStdtemplatesGetAPIRequest, resp *waybill.CainiaoCloudprintStdtemplatesGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
