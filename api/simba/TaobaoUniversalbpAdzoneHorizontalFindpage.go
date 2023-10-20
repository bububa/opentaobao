package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpadzonehorizontalfindpage 查看资源包列表
// taobao.universalbp.adzone.horizontal.findpage
//
// 查看已存在的计划上设置的资源包列表
func Taobaouniversalbpadzonehorizontalfindpage(clt *core.SDKClient, req *simba.TaobaouniversalbpadzonehorizontalfindpageAPIRequest, session string) (*simba.TaobaouniversalbpadzonehorizontalfindpageAPIResponse, error) {
	var resp simba.TaobaouniversalbpadzonehorizontalfindpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
