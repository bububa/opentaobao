package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

/* TaobaoOpenimTribeExpel
OPENIM群踢出成员
taobao.openim.tribe.expel

OPENIM群踢出成员 */
func TaobaoOpenimTribeExpel(clt *core.SDKClient, req *openim.TaobaoOpenimTribeExpelAPIRequest, session string) (*openim.TaobaoOpenimTribeExpelAPIResponse, error) {
	var resp openim.TaobaoOpenimTribeExpelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
