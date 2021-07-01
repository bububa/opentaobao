package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpIndustryLogisticsSyncAPIRequest
物流状态同步 API请求
alibaba.ascp.industry.logistics.sync

履约物流状态同步 */
type AlibabaAscpIndustryLogisticsSyncAPIRequest struct {
	model.Params
	// 参数
	_param *LogisticsSyncSellerRequest
}

// New
