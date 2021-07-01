package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmGradeGetAPIRequest
获取等级列表 API请求
alibaba.westcrm.grade.get

获取会员卡等级列表 */
type AlibabaWestcrmGradeGetAPIRequest struct {
	model.Params
	// 园区id
	_campusId int64
}

// New
