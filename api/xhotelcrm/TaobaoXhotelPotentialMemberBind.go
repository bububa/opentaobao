package xhotelcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelcrm"
)

// Taobaoxhotelpotentialmemberbind 飞猪酒店商家会员绑定
// taobao.xhotel.potential.member.bind
//
// 支持互通商家发起会员绑定
func Taobaoxhotelpotentialmemberbind(clt *core.SDKClient, req *xhotelcrm.TaobaoxhotelpotentialmemberbindAPIRequest, session string) (*xhotelcrm.TaobaoxhotelpotentialmemberbindAPIResponse, error) {
	var resp xhotelcrm.TaobaoxhotelpotentialmemberbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
