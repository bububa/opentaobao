package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCbossWorkplatformWorkorderCreateAPIRequest
菜鸟工单创建接口 API请求
cainiao.cboss.workplatform.workorder.create

菜鸟工单创建接口，目前调用者ISV */
type CainiaoCbossWorkplatformWorkorderCreateAPIRequest struct {
	model.Params
	// 工单类型
	_workOrderType string
	// 业务类型
	_bizType string
	// 工单创建备注
	_memo string
	// 货主商家用户id
	_memberId string
	// 货主用户角色
	_memberRole string
	// 创建者淘宝id（区分子账号）
	_creatorId string
	// 创建者角色
	_creatorRole string
	// 外部业务系统主键
	_bizEntityValue string
	// 店铺用户id
	_shopUserId string
	// 交易订单id
	_tradeId string
	// 工单来源
	_source string
	// 来源签名，用于唯一区分不同的来源方
	_sourceSign string
	// 运单号
	_mailNo string
	// 扩展字段
	_features string
	// 凭证地址列表
	_attachPathList []string
}

// New
