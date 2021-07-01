package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelBaseinfoCitiesGetAPIRequest
【API3.0】度假线路商品发布时基础信息获取接口：地址数据查询 API请求
taobao.alitrip.travel.baseinfo.cities.get

旅行度假新商品发布时可用的扩展接口，用于获取可用的出发地或目的地城市列表。 */
type TaobaoAlitripTravelBaseinfoCitiesGetAPIRequest struct {
	model.Params
	// 1-获取目的地城市列表 2-获取出发地城市列表
	_iocType int64
	// 1-境内跟团游 2-境内自由行 3-出境跟团游 4-出境自由行 5-境外当地玩乐 6-国际邮轮 9-境内邮轮
	_catType int64
}

// New
