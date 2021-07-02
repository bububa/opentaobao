package normalvisa

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// TaobaoAlitripTravelNormalvisaGetcompany 获取物流公司信息
// taobao.alitrip.travel.normalvisa.getcompany
//
// 获取物流公司信息
func TaobaoAlitripTravelNormalvisaGetcompany(clt *core.SDKClient, req *normalvisa.TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest, session string) (*normalvisa.TaobaoAlitripTravelNormalvisaGetcompanyAPIResponse, error) {
	var resp normalvisa.TaobaoAlitripTravelNormalvisaGetcompanyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
