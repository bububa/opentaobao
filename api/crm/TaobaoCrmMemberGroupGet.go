package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// Taobaocrmmembergroupget 获取买家身上的标签
// taobao.crm.member.group.get
//
// 获取买家身上的标签，不返回标签的总人数
func Taobaocrmmembergroupget(clt *core.SDKClient, req *crm.TaobaocrmmembergroupgetAPIRequest, session string) (*crm.TaobaocrmmembergroupgetAPIResponse, error) {
	var resp crm.TaobaocrmmembergroupgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
