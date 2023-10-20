package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbhouseDelete 民宿门店删除接口
// taobao.xhotel.bnbhouse.delete
//
// 支持门店相关的门店删除，删除门店会级联删除门店下面的房源
func TaobaoXhotelBnbhouseDelete(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbhouseDeleteAPIRequest, session string) (*xhotelitem.TaobaoXhotelBnbhouseDeleteAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelBnbhouseDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
