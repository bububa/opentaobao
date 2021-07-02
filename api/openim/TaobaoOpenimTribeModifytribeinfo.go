package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeModifytribeinfo OPENIM群信息修改
// taobao.openim.tribe.modifytribeinfo
//
// OPENIM群信息修改
func TaobaoOpenimTribeModifytribeinfo(clt *core.SDKClient, req *openim.TaobaoOpenimTribeModifytribeinfoAPIRequest, session string) (*openim.TaobaoOpenimTribeModifytribeinfoAPIResponse, error) {
	var resp openim.TaobaoOpenimTribeModifytribeinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
