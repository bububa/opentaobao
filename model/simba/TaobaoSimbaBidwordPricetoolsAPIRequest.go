package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaBidwordPricetoolsAPIRequest
关键词出价指导工具（新） API请求
taobao.simba.bidword.pricetools

关键词出价指导工具（新） */
type TaobaoSimbaBidwordPricetoolsAPIRequest struct {
	model.Params
	// 关键词id
	_bidwordId int64
	// 出价目标 ，1：争取排名；2：提升展现；3：提示点击；4：提升转化
	_type int64
	// 区分渠道 ，计算机：PC，无线 ：WL
	_trafficType string
	// 推广单元id
	_adgroupId int64
}

// New
