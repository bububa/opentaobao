package larkiot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLarkPosBasedataGetworkstationAPIRequest
根据影城id工作站和macId获取工作站 API请求
taobao.lark.pos.basedata.getworkstation

获取单独工作站 */
type TaobaoLarkPosBasedataGetworkstationAPIRequest struct {
	model.Params
	// 影城cinemaLinkId
	_cinemaLinkId string
	// 终端编码
	_posCode string
}

// New
