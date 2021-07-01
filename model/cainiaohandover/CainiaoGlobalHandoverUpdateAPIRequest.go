package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGlobalHandoverUpdateAPIRequest
修改交接单 API请求
cainiao.global.handover.update

提供给ISV通过该接口修改交接单 */
type CainiaoGlobalHandoverUpdateAPIRequest struct {
	model.Params
	// 用户信息
	_userInfo *UserInfoDto
	// 重量单位，克:g, 千克:kg，默认g
	_weightUnit string
	// 大包重量
	_weight int64
	// 交接单id
	_handoverOrderId int64
	// 大包备注
	_remark string
	// 退件信息
	_returnInfo *ReturnerDto
	// 揽收信息
	_pickupInfo *PickupDto
	// 交接单类型，菜鸟揽收(cainiao_pickup)或自寄(self_post)，默认菜鸟揽收
	_type string
	// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
	// 要创建交接单的小包编码集合，数量上限200
	_orderCodeList []string
	// 多语言
	_locale string
}

// New
