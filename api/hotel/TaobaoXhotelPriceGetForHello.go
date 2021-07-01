package hotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotel"
)

/* TaobaoXhotelPriceGetForHello
哈罗合作方获取酒店库存报价
taobao.xhotel.price.get.for.hello

哈罗合作项目，供哈罗合作方按需查询已开通城市下的标准酒店下指定时间段内的库存报价信息；在用户登录方面，返回结果不涉及用户个人信息，不涉及商家信息；仅根据不同用户，查询对应会员等级后，返回不同报价； */
func TaobaoXhotelPriceGetForHello(clt *core.SDKClient, req *hotel.TaobaoXhotelPriceGetForHelloAPIRequest, session string) (*hotel.TaobaoXhotelPriceGetForHelloAPIResponse, error) {
	var resp hotel.TaobaoXhotelPriceGetForHelloAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
