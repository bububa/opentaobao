package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojushitajdpuseradd 添加数据推送用户
// taobao.jushita.jdp.user.add
//
// 提供给接入数据推送的应用添加数据推送服务的用户
func Taobaojushitajdpuseradd(clt *core.SDKClient, req *jst.TaobaojushitajdpuseraddAPIRequest, session string) (*jst.TaobaojushitajdpuseraddAPIResponse, error) {
	var resp jst.TaobaojushitajdpuseraddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
