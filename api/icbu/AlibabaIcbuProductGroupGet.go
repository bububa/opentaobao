package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbuproductgroupget 分组信息获取
// alibaba.icbu.product.group.get
//
// 分组信息获取
func Alibabaicbuproductgroupget(clt *core.SDKClient, req *icbu.AlibabaicbuproductgroupgetAPIRequest, session string) (*icbu.AlibabaicbuproductgroupgetAPIResponse, error) {
	var resp icbu.AlibabaicbuproductgroupgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
