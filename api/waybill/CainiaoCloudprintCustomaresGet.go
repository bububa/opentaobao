package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoCloudprintCustomaresGet 获取商家的自定义区模板信息
// cainiao.cloudprint.customares.get
//
// 供isv使用，获取商家的自定义区的模板信息
func CainiaoCloudprintCustomaresGet(clt *core.SDKClient, req *waybill.CainiaoCloudprintCustomaresGetAPIRequest, resp *waybill.CainiaoCloudprintCustomaresGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
