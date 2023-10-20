package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojdshluserupdate 订单全链路用户信息修改
// taobao.jds.hluser.update
//
// 订单全链路用户信息修改，比如是否开放买家端展示
func Taobaojdshluserupdate(clt *core.SDKClient, req *jst.TaobaojdshluserupdateAPIRequest, session string) (*jst.TaobaojdshluserupdateAPIResponse, error) {
	var resp jst.TaobaojdshluserupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
