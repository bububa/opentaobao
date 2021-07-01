package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSeakingAuthmachineapiAPIRequest
机翻Api授权 API请求
alibaba.seaking.authmachineapi

机翻Api授权 */
type AlibabaSeakingAuthmachineapiAPIRequest struct {
	model.Params
	// erp名称
	_identifyType string
	// erp用户id
	_identifier string
	// 店铺所在平台
	_subIdentifyType string
	// 店铺id(ae为cn开头的店铺id, lazada为邮箱)
	_subIdentifier string
}

// New
