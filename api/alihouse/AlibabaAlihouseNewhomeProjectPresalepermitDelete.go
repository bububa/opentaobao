package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectPresalepermitDelete 删除楼盘预售证
// alibaba.alihouse.newhome.project.presalepermit.delete
//
// 删除楼盘预售证信息
func AlibabaAlihouseNewhomeProjectPresalepermitDelete(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
