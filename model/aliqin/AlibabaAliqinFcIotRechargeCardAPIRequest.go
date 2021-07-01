package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotRechargeCardAPIRequest
按终端号订购增值业务 API请求
alibaba.aliqin.fc.iot.rechargeCard

按终端号订购增值业务 */
type AlibabaAliqinFcIotRechargeCardAPIRequest struct {
	model.Params
	// 外部计费号类型：写‘ICCID’
	_billSource string
	// iccid的值
	_billReal string
	// 流量包offerId
	_offerId string
	// 外部id,用来做幂等
	_outRechargeId string
	// ICCID
	_iccid string
	// 生效时间,1000,立即生效; AUTO_ORD,下周期自动续订
	_effCode string
	// yyyy-MM-dd HH:mm:ss
	_effTime string
}

// New
