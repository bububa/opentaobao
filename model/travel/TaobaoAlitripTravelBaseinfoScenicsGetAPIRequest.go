package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest
【API3.0】基础信息获取接口：景点数据查询 API请求
taobao.alitrip.travel.baseinfo.scenics.get

商品发布辅助接口，用于飞猪度假或门票商品发布时 获取可用的景点（及景点下收费项目）信息列表。 */
type TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest struct {
	model.Params
	// 城市名称
	_city string
	// 景点简称
	_scenic string
	// 景点id，可选。若传了该值，则查询结果中只会保留该id的景点，其余景点信息将被过滤
	_scenicId int64
}

// New
