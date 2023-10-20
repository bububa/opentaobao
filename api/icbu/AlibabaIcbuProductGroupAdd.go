package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbuproductgroupadd 增加商品分组
// alibaba.icbu.product.group.add
//
// 增加商品分组
func Alibabaicbuproductgroupadd(clt *core.SDKClient, req *icbu.AlibabaicbuproductgroupaddAPIRequest, session string) (*icbu.AlibabaicbuproductgroupaddAPIResponse, error) {
	var resp icbu.AlibabaicbuproductgroupaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
