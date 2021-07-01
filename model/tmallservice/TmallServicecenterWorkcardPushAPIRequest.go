package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardPushAPIRequest
推送服务工单信息 API请求
tmall.servicecenter.workcard.push

服务商家推送工单信息到天猫。 */
type TmallServicecenterWorkcardPushAPIRequest struct {
	model.Params
	// 属性列表。使用半角分号隔开,字符串前后都需要有半角分号
	_attributes string
	// 描述
	_desc string
	// 淘宝交易订单号
	_bizOrderId int64
	// 服务预约安装时间
	_serviceReserveTime string
	// 服务预约安装地址。四级地址与街道地址用空格隔开
	_serviceReserveAddress string
	// 0=初始化, 3=授理， 10=拒绝 ，4=执行 ，5=成功，11=失败
	_status string
}

// New
