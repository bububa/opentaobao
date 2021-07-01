package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSscSupplyplatformServiceworkerSaveAPIRequest
服务商绑定工人 API请求
alibaba.ssc.supplyplatform.serviceworker.save

服务商将上传工人与服务商自己建立关系，需要将工人的服务区域和住址回传 */
type AlibabaSscSupplyplatformServiceworkerSaveAPIRequest struct {
	model.Params
	// 工人保存参数
	_workerSaveForTopReqDto *WorkerSaveForTopReqDto
}

// New
