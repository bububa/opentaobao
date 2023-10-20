package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoCloudprintIsvtemplatesGet 获取商家使用的标准模板
// cainiao.cloudprint.isvtemplates.get
//
// 获取商家使用的标准模板
func CainiaoCloudprintIsvtemplatesGet(clt *core.SDKClient, req *waybill.CainiaoCloudprintIsvtemplatesGetAPIRequest, resp *waybill.CainiaoCloudprintIsvtemplatesGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
