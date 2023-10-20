package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadgroupdeleteadgroupbatch 删除推广单元
// alibaba.scbp.ad.group.delete.ad.group.batch
//
// 删除推广单元
func Alibabascbpadgroupdeleteadgroupbatch(clt *core.SDKClient, req *scbp.AlibabascbpadgroupdeleteadgroupbatchAPIRequest, session string) (*scbp.AlibabascbpadgroupdeleteadgroupbatchAPIResponse, error) {
	var resp scbp.AlibabascbpadgroupdeleteadgroupbatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
