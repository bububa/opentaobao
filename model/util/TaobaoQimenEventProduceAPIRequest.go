package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenEventProduceAPIRequest
发出奇门事件 API请求
taobao.qimen.event.produce

当订单被处理时，用于通知奇门系统。 */
type TaobaoQimenEventProduceAPIRequest struct {
	model.Params
	// 事件状态，如QIMEN_ERP_TRANSFER，QIMEN_ERP_CHECK
	_status string
	// 淘宝订单号
	_tid string
	// JSON格式扩展字段
	_ext string
	// 商家平台编码.MAIN:官方渠道,JD:京东,DD:当当,PP:拍拍,YX:易讯,EBAY:ebay,AMAZON:亚马逊,SN:苏宁,GM:国美,WPH:唯品会,JM:聚美,MGJ:蘑菇街,YT:银泰,YHD:1号店,1688:1688,POS:POS门店,OTHER:其他
	_platform string
	// 订单创建时间,数字
	_create int64
	// 外部商家名称。必须同时填写platform
	_nick string
}

// New
