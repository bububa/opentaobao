package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSubwayCiaGetAPIRequest
查询单元智能出价信息 API请求
taobao.subway.cia.get

查询单元智能出价信息 */
type TaobaoSubwayCiaGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// New
