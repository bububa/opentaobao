package alihealthlab

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthlab"
)

// AlibabaAlihealthLabItemTbitemskuRelationSync 阿里健康检验检测业务，检验检测项目淘宝商品SKU关系同步
// alibaba.alihealth.lab.item.tbitemsku.relation.sync
//
// 阿里健康检验检测业务，检验检测项目淘宝商品SKU关系同步
func AlibabaAlihealthLabItemTbitemskuRelationSync(clt *core.SDKClient, req *alihealthlab.AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest, resp *alihealthlab.AlibabaAlihealthLabItemTbitemskuRelationSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
