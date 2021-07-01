package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripItFareDeleteAPIRequest
【国际机票自有政策】单条删除 API请求
taobao.alitrip.it.fare.delete

自有政策删除接口，可以根据fareId或outId删除，根据outId删除时，如果outId不唯一，返回失败 */
type TaobaoAlitripItFareDeleteAPIRequest struct {
	model.Params
	// json格式的字符串，扩展属性，预留
	_extendAttributes string
	// 运价id，单条新增成功时返回运价id，fareId和outId必填一个，fareId优先
	_fareId int64
	// 外部id，为新增时请求参数中的外部政策id
	_outId string
}

// New
