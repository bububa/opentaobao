package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIRequest
智能硬件互动云货架批量数据回流 API请求
taobao.smartstore.device.iashelf.batch.feedback

智慧门店互动云货架设备批量回流，
只能回流单个设备的批量业务数据
规则：
1. 回流的设备属于当前授权的用户
2. 回流的设备属于当前应用添加 */
type TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIRequest struct {
	model.Params
	// 硬件CODE
	_deviceCode string
	// 回流数据数组，一次最多100条
	_datas []DeviceBizDataDo
}

// New
