package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaoeticketmerchanttbmaget 码商查询淘宝码接口
// taobao.eticket.merchant.tbma.get
//
// 码商查询淘宝码接口
func Taobaoeticketmerchanttbmaget(clt *core.SDKClient, req *eticket.TaobaoeticketmerchanttbmagetAPIRequest, session string) (*eticket.TaobaoeticketmerchanttbmagetAPIResponse, error) {
	var resp eticket.TaobaoeticketmerchanttbmagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
