package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpAicSupplierAicinventoryPublishAPIRequest
商家仓操作aic库存发布服务 API请求
alibaba.ascp.aic.supplier.aicinventory.publish

商家调用这个接口来发布增加库存数据 */
type AlibabaAscpAicSupplierAicinventoryPublishAPIRequest struct {
	model.Params
	// 库存发布请求参数
	_aicInventoryPublishRequest *Aicinventorypublishrequest
}

// New
