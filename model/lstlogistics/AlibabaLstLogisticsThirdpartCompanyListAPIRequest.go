package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstLogisticsThirdpartCompanyListAPIRequest
供应商-异云-第三方物流公司列表 API请求
alibaba.lst.logistics.thirdpart.company.list

异地云仓发货时，需填写的第三方物流公司列表 */
type AlibabaLstLogisticsThirdpartCompanyListAPIRequest struct {
	model.Params
	// 入参
	_query *LstLogisticsCompanyQuery
}

// New
