package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSerchcrowdGetAPIRequest
根据推广单元id获取搜索溢价人群 API请求
taobao.simba.serchcrowd.get

根据推广单元id获取搜索溢价人群 */
type TaobaoSimbaSerchcrowdGetAPIRequest struct {
	model.Params
	// 被操作者的淘宝昵称
	_nick string
	// 推广单元id
	_adgroupId int64
}

// New
