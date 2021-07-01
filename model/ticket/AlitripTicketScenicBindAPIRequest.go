package ticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTicketScenicBindAPIRequest
【门票API2.0】门票景点绑定接口 API请求
alitrip.ticket.scenic.bind

门票景点绑定接口，用于建立阿里标准景点id与商家系统景点id的映射关系。该接口同时支持新建和修改映射关系，当用户没有为ali_scenic_id建立过映射关系时，则判断为新建映射关系，否则为修改。可以通过设置update_out_scenic_id来修改ali_scenic_id与out_scenic_id的映射关系。 */
type AlitripTicketScenicBindAPIRequest struct {
	model.Params
	// 必填，阿里旅行对应的景点编码
	_aliScenicId int64
	// 商户景点城市
	_city string
	// 商户景点地址
	_address string
	// 商户景点名称
	_outScenicName string
	// 商户景点省份
	_province string
	// 必填，商户系统中的景点编码，用于与ali_scenic_id建立映射关系
	_outScenicId string
	// 可选，如果需要更新out_scenic_id与ali_scenic_id的映射关系时 需要填写
	_updateOutScenicId string
}

// New
