package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbakeywordsbykeywordidsget 根据一个关键词Id列表取得一组关键词
// taobao.simba.keywordsbykeywordids.get
//
// 根据一个关键词Id列表取得一组关键词
func Taobaosimbakeywordsbykeywordidsget(clt *core.SDKClient, req *simba.TaobaosimbakeywordsbykeywordidsgetAPIRequest, session string) (*simba.TaobaosimbakeywordsbykeywordidsgetAPIResponse, error) {
	var resp simba.TaobaosimbakeywordsbykeywordidsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
