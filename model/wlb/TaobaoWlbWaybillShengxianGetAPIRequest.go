package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWaybillShengxianGetAPIRequest
商家获取生鲜电子面单号 API请求
taobao.wlb.waybill.shengxian.get

商家通过交易订单号获取电子面单接口 */
type TaobaoWlbWaybillShengxianGetAPIRequest struct {
	model.Params
	// 物流服务方代码，生鲜配送：YXSR
	_bizCode string
	// 物流服务类型。冷链：602，常温：502
	_deliveryType string
	// 商家淘宝地址信息ID
	_senderAddressId string
	// 仓库的服务代码标示，代码一个仓库的实体。(可以通过taobao.wlb.stores.baseinfo.get接口查询)
	_serviceCode string
	// 订单渠道： 淘宝平台订单："TB"; 天猫平台订单："TM"; 京东："JD"; 拍拍："PP"; 易迅平台订单："YX"; 一号店平台订单："YHD"; 当当网平台订单："DD"; EBAY："EBAY"; QQ网购："QQ"; 苏宁："SN"; 国美："GM"; 唯品会："WPH"; 聚美："Jm"; 乐峰："LF"; 蘑菇街："MGJ"; 聚尚："JS"; 银泰："YT"; VANCL："VANCL"; 邮乐："YL"; 优购："YG"; 拍鞋："PX"; 1688平台："1688";
	_orderChannelsType string
	// 交易号，传入交易号不能存在拆单场景。
	_tradeId string
	// 预留扩展字段
	_feature string
}

// New
