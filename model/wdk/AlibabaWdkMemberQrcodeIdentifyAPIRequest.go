package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMemberQrcodeIdentifyAPIRequest
扫码识别会员接口 API请求
alibaba.wdk.member.qrcode.identify

根据用户输入的付款码（支付宝、盒马、淘宝）、商家等信息，查询当前用户的基本信息及对应会员卡信息 */
type AlibabaWdkMemberQrcodeIdentifyAPIRequest struct {
	model.Params
	// 付款码
	_qrCode string
}

// New
