package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojushitajdpuserdelete 删除数据推送用户
// taobao.jushita.jdp.user.delete
//
// 删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。
func Taobaojushitajdpuserdelete(clt *core.SDKClient, req *jst.TaobaojushitajdpuserdeleteAPIRequest, session string) (*jst.TaobaojushitajdpuserdeleteAPIResponse, error) {
	var resp jst.TaobaojushitajdpuserdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
