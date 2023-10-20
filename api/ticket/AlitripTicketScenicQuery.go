package ticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ticket"
)

// Alitripticketscenicquery 【门票API2.0】卖家已发布门票商品列表查询接口（根据景点维度查询）
// alitrip.ticket.scenic.query
//
// 查询卖家已发布过的门票商品列表，根据景点维度聚合查询。如果卖家在该景点下未发布过任何商品，则查询不到数据！
func Alitripticketscenicquery(clt *core.SDKClient, req *ticket.AlitripticketscenicqueryAPIRequest, session string) (*ticket.AlitripticketscenicqueryAPIResponse, error) {
	var resp ticket.AlitripticketscenicqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
