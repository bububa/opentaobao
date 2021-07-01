package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrFulfillLogisticsConsignAPIRequest
同城配门店备货通知 API请求
tmall.nr.fulfill.logistics.consign

同城配业务备货通知，商家告诉平台门店的货已经准备好，可以发货了； */
type TmallNrFulfillLogisticsConsignAPIRequest struct {
	model.Params
	// 入参对象
	_param0 *NrStoreGoodsReadyReqDto
}

// New
