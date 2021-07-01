package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbCrossborderWaybillGetAPIRequest
集货商家pdf和云打印面单获取，pdf需要配置白名单 API请求
taobao.wlb.crossborder.waybill.get

【TOF】欧洲供应商PDF格式电子面单渲染下发
 需求链接：https://aone.alibaba-inc.com/req/21210808 */
type TaobaoWlbCrossborderWaybillGetAPIRequest struct {
	model.Params
	// 菜鸟物流单号
	_orderCode string
}

// New
