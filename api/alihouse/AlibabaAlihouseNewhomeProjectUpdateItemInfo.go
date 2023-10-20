package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectUpdateItemInfo 更新楼盘商品信息
// alibaba.alihouse.newhome.project.update.item.info
//
// 更新楼盘商品信息
func AlibabaAlihouseNewhomeProjectUpdateItemInfo(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectUpdateItemInfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
