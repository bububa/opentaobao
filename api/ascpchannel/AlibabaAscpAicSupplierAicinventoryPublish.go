package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpAicSupplierAicinventoryPublish 商家仓操作aic库存发布服务
// alibaba.ascp.aic.supplier.aicinventory.publish
//
// 商家调用这个接口来发布增加库存数据
func AlibabaAscpAicSupplierAicinventoryPublish(clt *core.SDKClient, req *ascpchannel.AlibabaAscpAicSupplierAicinventoryPublishAPIRequest, resp *ascpchannel.AlibabaAscpAicSupplierAicinventoryPublishAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
