package c2m

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSebpOrganizationGetorderinfoAPIRequest
淘小铺机构订单信息 API请求
taobao.sebp.organization.getorderinfo

淘小铺合作机构获取机构订单信息，用于机构结算使用 */
type TaobaoSebpOrganizationGetorderinfoAPIRequest struct {
	model.Params
	// null-请求所有，20200616-请求2020年6月16号的变更信息
	_modifyDate string
	// 第几页
	_pageNum int64
	// 查询实时数据时，必传，开始时间结束时间间隔不能超过4个小时
	_endTime string
	// 查询实时数据时，必传，开始时间不能早于2天前
	_startTime string
}

// New
