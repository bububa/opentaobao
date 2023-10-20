package normalvisa

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/normalvisa"
)

// TaobaoAlitripTravelNormalvisaGetdetail 获取单笔订单的详情
// taobao.alitrip.travel.normalvisa.getdetail
//
// 获取单笔签证的详细记录
func TaobaoAlitripTravelNormalvisaGetdetail(clt *core.SDKClient, req *normalvisa.TaobaoAlitripTravelNormalvisaGetdetailAPIRequest, session string) (*normalvisa.TaobaoAlitripTravelNormalvisaGetdetailAPIResponse, error) {
	var resp normalvisa.TaobaoAlitripTravelNormalvisaGetdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
