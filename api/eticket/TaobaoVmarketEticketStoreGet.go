package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaovmarketeticketstoreget 获取电子凭证预约门店信息
// taobao.vmarket.eticket.store.get
//
// 用于给外部商家查询电子凭证预约门店信息
func Taobaovmarketeticketstoreget(clt *core.SDKClient, req *eticket.TaobaovmarketeticketstoregetAPIRequest, session string) (*eticket.TaobaovmarketeticketstoregetAPIResponse, error) {
	var resp eticket.TaobaovmarketeticketstoregetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
