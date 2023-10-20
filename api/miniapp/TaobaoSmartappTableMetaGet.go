package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoSmartappTableMetaGet 智能应用服务登记工作表元数据查询
// taobao.smartapp.table.meta.get
//
// 智能应用服务登记工作表元数据查询
func TaobaoSmartappTableMetaGet(clt *core.SDKClient, req *miniapp.TaobaoSmartappTableMetaGetAPIRequest, resp *miniapp.TaobaoSmartappTableMetaGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
