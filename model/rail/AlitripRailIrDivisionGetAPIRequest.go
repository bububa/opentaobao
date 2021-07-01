package rail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripRailIrDivisionGetAPIRequest
国际火车票标准城市查询 API请求
alitrip.rail.ir.division.get

国际火车票提供给代理商用于查询标准城市信息，全部城市数据量209530条，含除中国大陆以外的全部海外区域。
代理商通过分页查询的方式，拉取飞猪平台方全部境外标准城市，用于自身城市与飞猪平台城市的映射。 */
type AlitripRailIrDivisionGetAPIRequest struct {
	model.Params
	// 代理商id
	_agentId int64
	// 层级，1洲，2是国家，3是省，4是市，5是区，6是街道/镇，7是村，8是逻辑行政区，境外火车票业务只需要市级别，传4就可以
	_level int64
	// 每页条数
	_pageSize int64
	// 页数，从1开始
	_pageIndex int64
}

// New
