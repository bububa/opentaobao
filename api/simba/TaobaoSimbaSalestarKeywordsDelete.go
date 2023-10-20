package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbasalestarkeywordsdelete 销量明星关键词删除
// taobao.simba.salestar.keywords.delete
//
// （新）关键词删除相关接口
func Taobaosimbasalestarkeywordsdelete(clt *core.SDKClient, req *simba.TaobaosimbasalestarkeywordsdeleteAPIRequest, session string) (*simba.TaobaosimbasalestarkeywordsdeleteAPIResponse, error) {
	var resp simba.TaobaosimbasalestarkeywordsdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
