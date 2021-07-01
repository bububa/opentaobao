package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrFulfillLogisticsSyncAPIRequest
同城配物流信息回传 API请求
tmall.nr.fulfill.logistics.sync

同城配业务物流信息回传，通过接口将物流信息同步给天猫 */
type TmallNrFulfillLogisticsSyncAPIRequest struct {
	model.Params
	// 物流回传参数
	_param0 *NrLogisticsInfoSynReqDto
}

// New
