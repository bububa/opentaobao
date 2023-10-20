package normalvisa

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// TaobaoAlitripTravelNormalvisaGet 获取签证记录
// taobao.alitrip.travel.normalvisa.get
//
// 用于获取普通签证的记录信息
func TaobaoAlitripTravelNormalvisaGet(clt *core.SDKClient, req *normalvisa.TaobaoAlitripTravelNormalvisaGetAPIRequest, resp *normalvisa.TaobaoAlitripTravelNormalvisaGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
