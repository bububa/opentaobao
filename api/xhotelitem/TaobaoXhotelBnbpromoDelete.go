package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbpromoDelete 民宿卖家活动删除
// taobao.xhotel.bnbpromo.delete
//
// 民宿删除营销活动
func TaobaoXhotelBnbpromoDelete(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbpromoDeleteAPIRequest, session string) (*xhotelitem.TaobaoXhotelBnbpromoDeleteAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelBnbpromoDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
