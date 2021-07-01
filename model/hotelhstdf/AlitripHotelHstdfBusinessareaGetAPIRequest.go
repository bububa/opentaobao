package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelHstdfBusinessareaGetAPIRequest
根据城市查询商圈 API请求
alitrip.hotel.hstdf.businessarea.get

根据cityId分页查询商圈信息 */
type AlitripHotelHstdfBusinessareaGetAPIRequest struct {
	model.Params
	// 请求参数封装
	_paramGetByTrdiDivisionIdParam *GetByTrdiDivisionIdParam
}

// New
