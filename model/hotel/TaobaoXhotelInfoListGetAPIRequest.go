package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelInfoListGetAPIRequest
酒店详细信息查询 API请求
taobao.xhotel.info.list.get

获取酒店详情信息 */
type TaobaoXhotelInfoListGetAPIRequest struct {
	model.Params
	// 城市code
	_cityCode int64
	// 分页参数：当前页数，从1开始计数。<br/>默认值：1
	_currentPage int64
	// 分页参数：每页酒店数量。最小值1，最大值为50。默认值：20
	_pageSize int64
	// pid
	_pid string
	// 标准酒店id，如果需要查询单条酒店的信息，需要传入此参数
	_shid int64
}

// New
