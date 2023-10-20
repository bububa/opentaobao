package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuPhotobankGroupList 图片银行分组信息获取
// alibaba.icbu.photobank.group.list
//
// 图片银行分组信息获取
func AlibabaIcbuPhotobankGroupList(clt *core.SDKClient, req *icbu.AlibabaIcbuPhotobankGroupListAPIRequest, session string) (*icbu.AlibabaIcbuPhotobankGroupListAPIResponse, error) {
	var resp icbu.AlibabaIcbuPhotobankGroupListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
