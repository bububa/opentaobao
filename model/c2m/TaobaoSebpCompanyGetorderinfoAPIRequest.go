package c2m

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSebpCompanyGetorderinfoAPIRequest
淘小铺公司订单信息 API请求
taobao.sebp.company.getorderinfo

淘小铺合作公司获取公司订单信息，用于公司结算使用 */
type TaobaoSebpCompanyGetorderinfoAPIRequest struct {
	model.Params
	// null-请求所有，20200616-请求2020年6月16号的变更信息
	_modifyDate string
	// 第几页
	_pageNum int64
}

// New
