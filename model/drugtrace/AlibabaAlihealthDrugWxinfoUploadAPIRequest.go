package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugWxinfoUploadAPIRequest
小程序数据回传 API请求
alibaba.alihealth.drug.wxinfo.upload

小程序数据回传 */
type AlibabaAlihealthDrugWxinfoUploadAPIRequest struct {
	model.Params
	// 用户信息
	_userInfo string
	// 店铺名称
	_shopInfo string
	// 售货员信息
	_salerInfo string
	// 渠道
	_isvChannel string
}

// New
