package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbuphotobankgrouplist 图片银行分组信息获取
// alibaba.icbu.photobank.group.list
//
// 图片银行分组信息获取
func Alibabaicbuphotobankgrouplist(clt *core.SDKClient, req *icbu.AlibabaicbuphotobankgrouplistAPIRequest, session string) (*icbu.AlibabaicbuphotobankgrouplistAPIResponse, error) {
	var resp icbu.AlibabaicbuphotobankgrouplistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
