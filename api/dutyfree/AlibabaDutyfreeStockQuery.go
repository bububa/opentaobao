package dutyfree

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dutyfree"
)

// Alibabadutyfreestockquery 对外库存查询接口
// alibaba.dutyfree.stock.query
//
// 对外部服务提供库存查询接口
func Alibabadutyfreestockquery(clt *core.SDKClient, req *dutyfree.AlibabadutyfreestockqueryAPIRequest, session string) (*dutyfree.AlibabadutyfreestockqueryAPIResponse, error) {
	var resp dutyfree.AlibabadutyfreestockqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
