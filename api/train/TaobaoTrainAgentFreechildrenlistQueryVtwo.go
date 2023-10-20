package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentfreechildrenlistqueryvtwo 免费儿童列表查询
// taobao.train.agent.freechildrenlist.query.vtwo
//
// 免费儿童列表查询
func Taobaotrainagentfreechildrenlistqueryvtwo(clt *core.SDKClient, req *train.TaobaotrainagentfreechildrenlistqueryvtwoAPIRequest, session string) (*train.TaobaotrainagentfreechildrenlistqueryvtwoAPIResponse, error) {
	var resp train.TaobaotrainagentfreechildrenlistqueryvtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
