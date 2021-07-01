package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGlobalHandoverSavedraftAPIRequest
创建交接单草稿 API请求
cainiao.global.handover.savedraft

提供给ISV通过该接口创建交接单草稿 */
type CainiaoGlobalHandoverSavedraftAPIRequest struct {
	model.Params
	// 用户信息
	_userInfo *UserInfoDto
	// 备注
	_remark string
	// 大包重量
	_weight int64
	// 重量单位，克:g, 千克:kg，默认g
	_weightUnit string
	// 揽收信息
	_pickupInfo *PickupDto
	// 退件信息
	_returnInfo *ReturnerDto
	// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
	// 需要组装大包的小包编码集合，最多限制200个小包
	_orderCodeList []string
	// 多语言
	_locale string
}

// New
