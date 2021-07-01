package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest
物流工单履约状态更新 API请求
tmall.servicecenter.workcard.logisticsorder.update

天猫寄送类服务对接外部物流服务商回传物流状态信息 */
type TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest struct {
	model.Params
	// 体积 单位 立方毫米
	_volume int64
	// 重量 单位 克
	_weight int64
	// 备注说明
	_comment string
	// 物流单号（展示给消费者）
	_expressCode string
	// 物流公司名词（展示给消费者）
	_expressCompany string
	// 小件员手机号码
	_courierMobile string
	// 小件员姓名
	_courierName string
	// 取件码
	_gotCode string
	// 物流订单号
	_logisticsOrderId int64
	// 金额 单位分
	_cost int64
	// 1、以下状态时必填： 包裹已揽收完成 2、字符串格式为：json 串 例子： [{ "name": "上衣", "count": 1 }, { "name": "裤子", "count": 2 }]
	_goodsInfo string
	// 物流创建 ：create 物流取消 ：cancel 分派小件员：assign 已经分派小件员: assigned 包裹上门揽收: pickup_door 包裹已揽收完成: pickup_finished 包裹派送中: dispatching 包裹已签收: signed
	_statusCode string
	// 物流子单号
	_subExpressCodes []string
	// 预计送达时间  dispatching节点时必填
	_deliveryTime string
	// 签收时间 signed节点时必填
	_signTime string
	// 揽收完成时间 pickup_finished节点时必填
	_pickupFinishTime string
	// 上门揽收时间 pickup_door节点时必填
	_pickupDoorTime string
}

// New
