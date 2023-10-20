package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadgroupcreateadgroupbatch 创建推广单元
// alibaba.scbp.ad.group.create.ad.group.batch
//
// 创建推广单元
func Alibabascbpadgroupcreateadgroupbatch(clt *core.SDKClient, req *scbp.AlibabascbpadgroupcreateadgroupbatchAPIRequest, session string) (*scbp.AlibabascbpadgroupcreateadgroupbatchAPIResponse, error) {
	var resp scbp.AlibabascbpadgroupcreateadgroupbatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
