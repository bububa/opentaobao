package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSubwayWordpackageGetAPIRequest
获取词包列表 API请求
taobao.subway.wordpackage.get

获取流量智选、捡漏词包等词包列表 */
type TaobaoSubwayWordpackageGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组id
	_adgroupId int64
}

// New
