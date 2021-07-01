package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripItFareGetAPIRequest
【国际机票自有政策】单条查询 API请求
taobao.alitrip.it.fare.get

通过此接口可以查询单条政策的详情，可以根据fareId或outId查询，用户outId查询时，如果outId不唯一，只返回最新添加的一条数据 */
type TaobaoAlitripItFareGetAPIRequest struct {
	model.Params
	// json格式的字符串，扩展属性，预留
	_extendAttributes string
	// 运价id，单条新增成功时返回运价id，fareId和outId必填一个，fareId优先
	_fareId int64
	// 外部id，为新增时请求参数中的外部政策id
	_outId string
}

// New
