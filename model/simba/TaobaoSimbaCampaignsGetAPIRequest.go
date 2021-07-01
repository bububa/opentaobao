package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCampaignsGetAPIRequest
取得一组推广计划 API请求
taobao.simba.campaigns.get

取得一个客户的推广计划； */
type TaobaoSimbaCampaignsGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 计划类型0位标准计划，16位销量明星计划
	_type int64
}

// New
