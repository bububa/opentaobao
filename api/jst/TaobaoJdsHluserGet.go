package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojdshluserget 订单全链路用户信息获取
// taobao.jds.hluser.get
//
// 订单全链路用户信息获取
func Taobaojdshluserget(clt *core.SDKClient, req *jst.TaobaojdshlusergetAPIRequest, session string) (*jst.TaobaojdshlusergetAPIResponse, error) {
	var resp jst.TaobaojdshlusergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
