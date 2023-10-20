package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpproductgroupget 查询指定产品分组的下一层子分组
// alibaba.scbp.product.group.get
//
// 查询指定产品分组的下一层子分组
func Alibabascbpproductgroupget(clt *core.SDKClient, req *scbp.AlibabascbpproductgroupgetAPIRequest, session string) (*scbp.AlibabascbpproductgroupgetAPIResponse, error) {
	var resp scbp.AlibabascbpproductgroupgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
