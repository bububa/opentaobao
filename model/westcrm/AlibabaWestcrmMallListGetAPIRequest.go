package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmMallListGetAPIRequest
获取商场列表 API请求
alibaba.westcrm.mall.list.get

根据园区id获取商场列表 */
type AlibabaWestcrmMallListGetAPIRequest struct {
	model.Params
	// 园区id
	_campusId int64
}

// New
