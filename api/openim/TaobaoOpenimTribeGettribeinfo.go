package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

/* TaobaoOpenimTribeGettribeinfo
获取群信息
taobao.openim.tribe.gettribeinfo

获取群信息 */
func TaobaoOpenimTribeGettribeinfo(clt *core.SDKClient, req *openim.TaobaoOpenimTribeGettribeinfoAPIRequest, session string) (*openim.TaobaoOpenimTribeGettribeinfoAPIResponse, error) {
	var resp openim.TaobaoOpenimTribeGettribeinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
