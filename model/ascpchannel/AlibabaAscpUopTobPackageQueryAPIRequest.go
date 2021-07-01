package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopTobPackageQueryAPIRequest
B2B包裹查询接口 API请求
alibaba.ascp.uop.tob.package.query

供应链中台TOB包裹查询接口 */
type AlibabaAscpUopTobPackageQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_packageQueryRequest *Packagequeryrequest
}

// New
