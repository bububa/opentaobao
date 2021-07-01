package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceDeviceOrderUpdateAPIRequest
回传/更新设备订购单 API请求
alibaba.einvoice.device.order.update

更新设备订购单，同步税控设备信息 */
type AlibabaEinvoiceDeviceOrderUpdateAPIRequest struct {
	model.Params
	// 订购单工单事件：  deploy_finish: 设备就绪，部署完成  isv_reject: 服务商驳回订购单
	_action string
	// 税控设备ID
	_deviceId string
	// 拓展字段。  ①当action=deploy_finish时，拓展字段中必须包含：  serv_start_time: 服务有效周期-起始时间  serv_end_time: 服务有效周期-结束时间  时间格式：yyyy-MM-dd HH:mm:ss  ②当action=isv_reject时，拓展字段中必须包含：  message: 驳回原因
	_extJson string
	// 订购开通单ID
	_flowId string
	// 税号
	_payeeRegisterNo string
}

// New
