package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbakeywordsdelete 删除一批关键词
// taobao.simba.keywords.delete
//
// 删除一批关键词
func Taobaosimbakeywordsdelete(clt *core.SDKClient, req *simba.TaobaosimbakeywordsdeleteAPIRequest, session string) (*simba.TaobaosimbakeywordsdeleteAPIResponse, error) {
	var resp simba.TaobaosimbakeywordsdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
