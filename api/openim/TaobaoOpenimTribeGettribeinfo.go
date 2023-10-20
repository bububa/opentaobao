package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimtribegettribeinfo 获取群信息
// taobao.openim.tribe.gettribeinfo
//
// 获取群信息
func Taobaoopenimtribegettribeinfo(clt *core.SDKClient, req *openim.TaobaoopenimtribegettribeinfoAPIRequest, session string) (*openim.TaobaoopenimtribegettribeinfoAPIResponse, error) {
	var resp openim.TaobaoopenimtribegettribeinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
