package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceFloorGetbyidAPIRequest
根据id获取楼层 API请求
alibaba.campus.space.floor.getbyid

根据id获取楼层 */
type AlibabaCampusSpaceFloorGetbyidAPIRequest struct {
	model.Params
	// 环境上下文
	_context *WorkBenchContext
	// 楼层id
	_id int64
}

// New
