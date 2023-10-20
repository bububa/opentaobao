package jms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jms"
)

// Taobaojushitajmsuserdelete 删除ONS消息同步用户
// taobao.jushita.jms.user.delete
//
// 删除ONS消息同步用户，删除后用户的消息将不会推送到聚石塔的ONS中
func Taobaojushitajmsuserdelete(clt *core.SDKClient, req *jms.TaobaojushitajmsuserdeleteAPIRequest, session string) (*jms.TaobaojushitajmsuserdeleteAPIResponse, error) {
	var resp jms.TaobaojushitajmsuserdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
