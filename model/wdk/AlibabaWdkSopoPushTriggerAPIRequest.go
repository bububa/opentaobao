package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSopoPushTriggerAPIRequest
猫超共享库存寄售sopo推送触发 API请求
alibaba.wdk.sopo.push.trigger

猫超共享库存寄售sopo触发推送给商家 */
type AlibabaWdkSopoPushTriggerAPIRequest struct {
	model.Params
	// 系统自动生成
	_wdkOpenPushSoPoRequest *WdkOpenPushSoPoRequest
}

// New
