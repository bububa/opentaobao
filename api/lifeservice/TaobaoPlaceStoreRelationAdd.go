package lifeservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lifeservice"
)

// Taobaoplacestorerelationadd 门店关系新增
// taobao.place.store.relation.add
//
// 新增授权用户的门店关系信息
func Taobaoplacestorerelationadd(clt *core.SDKClient, req *lifeservice.TaobaoplacestorerelationaddAPIRequest, session string) (*lifeservice.TaobaoplacestorerelationaddAPIResponse, error) {
	var resp lifeservice.TaobaoplacestorerelationaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
