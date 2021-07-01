package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripItFareBatchdeleteAPIRequest
【国际机票自有政策】批量删除 API请求
taobao.alitrip.it.fare.batchdelete

批量删除自有政策，单次删除最大5万，大于5万时候提示失败，需要缩小删除条件。此接口同步返回任务id，异步执行操作。每个接入方最多同时只能有10个处理中的任务，超过后直接返回失败。 */
type TaobaoAlitripItFareBatchdeleteAPIRequest struct {
	model.Params
	// 航空公司
	_airline string
	// 到达城市 可传多个 AND关系
	_arrCity string
	// 舱位 可传多个 或者的关系
	_cabin string
	// 是否能够混舱
	_canRt bool
	// 到达城市 可传多个 AND关系
	_depCity string
	// 最晚修改时间
	_endModifyDate string
	// 文件编号
	_fileCode string
	// 维护方式，可选值（UI：后台界面录入；EXCEL：后台excel批量导入；API：top接口添加）
	_operateSource string
	// 外部政策id
	_outId string
	// 最早修改时间
	_startModifyDate string
	// 去程适用开始日期
	_startRestrictGoDate string
	// 去程适用结束日期
	_endRestrictGoDate string
	// 运价类型，1单程 2往返
	_fareType int64
	// 0：未发布 1：已发布 2：已过期。不传的话，默认只能删除未发布和已过期的数据
	_statusList []int64
	// json格式的字符串，扩展属性，预留
	_extendAttributes string
}

// New
