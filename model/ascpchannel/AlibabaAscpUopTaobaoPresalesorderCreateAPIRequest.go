package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest
预售商家仓接单 API请求
alibaba.ascp.uop.taobao.presalesorder.create

预售商家仓接单 */
type AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest struct {
	model.Params
	// 预售商家仓接单对象
	_presalesOrderCreateRequest *PresalesordercreaterequestTest
}

// New
