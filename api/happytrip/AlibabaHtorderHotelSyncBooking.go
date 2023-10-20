package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahtorderhotelsyncbooking 未来酒店亲橙客栈预订信息同步
// alibaba.htorder.hotel.sync.booking
//
// 未来酒店亲橙客栈预订信息同步
func Alibabahtorderhotelsyncbooking(clt *core.SDKClient, req *happytrip.AlibabahtorderhotelsyncbookingAPIRequest, session string) (*happytrip.AlibabahtorderhotelsyncbookingAPIResponse, error) {
	var resp happytrip.AlibabahtorderhotelsyncbookingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
