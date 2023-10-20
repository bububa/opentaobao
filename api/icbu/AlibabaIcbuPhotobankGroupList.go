package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuPhotobankGroupList 图片银行分组信息获取
// alibaba.icbu.photobank.group.list
//
// 图片银行分组信息获取
func AlibabaIcbuPhotobankGroupList(clt *core.SDKClient, req *icbu.AlibabaIcbuPhotobankGroupListAPIRequest, resp *icbu.AlibabaIcbuPhotobankGroupListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
