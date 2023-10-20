package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadgroupupdateadgroupbatch 修改推广单元
// alibaba.scbp.ad.group.update.ad.group.batch
//
// 修改推广单元
func Alibabascbpadgroupupdateadgroupbatch(clt *core.SDKClient, req *scbp.AlibabascbpadgroupupdateadgroupbatchAPIRequest, session string) (*scbp.AlibabascbpadgroupupdateadgroupbatchAPIResponse, error) {
	var resp scbp.AlibabascbpadgroupupdateadgroupbatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
