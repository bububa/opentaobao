package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCardBindcardAPIRequest
绑定物理卡 API请求
alibaba.alsc.crm.card.bindcard

将会员卡和实例物理卡绑定在一起 */
type AlibabaAlscCrmCardBindcardAPIRequest struct {
	model.Params
	// 请求参数
	_paramBindPhysicalCardOpenReq *BindPhysicalCardOpenReq
}

// New
