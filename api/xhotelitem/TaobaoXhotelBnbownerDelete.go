package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbownerDelete 民宿房东删除接口
// taobao.xhotel.bnbowner.delete
//
// 民宿房东删除接口，删除房东后，对应的门店及房源会同步删除，请谨慎使用
func TaobaoXhotelBnbownerDelete(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbownerDeleteAPIRequest, session string) (*xhotelitem.TaobaoXhotelBnbownerDeleteAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelBnbownerDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
