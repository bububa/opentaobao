package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// Taobaotmcgroupdelete 删除指定的分组或分组下的用户
// taobao.tmc.group.delete
//
// 删除指定的分组或分组下的用户，授权消息使用
func Taobaotmcgroupdelete(clt *core.SDKClient, req *tmc.TaobaotmcgroupdeleteAPIRequest, session string) (*tmc.TaobaotmcgroupdeleteAPIResponse, error) {
	var resp tmc.TaobaotmcgroupdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
