package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminuserorderlist 获取用户订单列表
// yunos.tvpubadmin.user.orderlist
//
// 获取用户订单列表
func Yunostvpubadminuserorderlist(clt *core.SDKClient, req *tvupadmin.YunostvpubadminuserorderlistAPIRequest, session string) (*tvupadmin.YunostvpubadminuserorderlistAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminuserorderlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
