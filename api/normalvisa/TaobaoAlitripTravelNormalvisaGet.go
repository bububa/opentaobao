package normalvisa

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// TaobaoAlitripTravelNormalvisaGet 获取签证记录
// taobao.alitrip.travel.normalvisa.get
//
// 用于获取普通签证的记录信息
func TaobaoAlitripTravelNormalvisaGet(clt *core.SDKClient, req *normalvisa.TaobaoAlitripTravelNormalvisaGetAPIRequest, session string) (*normalvisa.TaobaoAlitripTravelNormalvisaGetAPIResponse, error) {
	var resp normalvisa.TaobaoAlitripTravelNormalvisaGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
