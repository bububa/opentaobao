package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojstminiappcrowdcreate 小程序活动创建
// taobao.jst.miniapp.crowd.create
//
// 小程序活动创建
func Taobaojstminiappcrowdcreate(clt *core.SDKClient, req *jst.TaobaojstminiappcrowdcreateAPIRequest, session string) (*jst.TaobaojstminiappcrowdcreateAPIResponse, error) {
	var resp jst.TaobaojstminiappcrowdcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
