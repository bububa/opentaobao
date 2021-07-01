package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopTaobaoWaybillQuerycpAPIRequest
查询电子面单开放的CP列表 API请求
alibaba.ascp.uop.taobao.waybill.querycp

查询电子面单开放的CP列表 */
type AlibabaAscpUopTaobaoWaybillQuerycpAPIRequest struct {
	model.Params
	// 系统自动生成
	_queryCpRequest *Querycprequest
}

// New
