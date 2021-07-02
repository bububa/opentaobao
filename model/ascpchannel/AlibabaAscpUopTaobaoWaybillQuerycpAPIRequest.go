package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopTaobaoWaybillQuerycpAPIRequest 查询电子面单开放的CP列表 API请求
// alibaba.ascp.uop.taobao.waybill.querycp
//
// 查询电子面单开放的CP列表
type AlibabaAscpUopTaobaoWaybillQuerycpAPIRequest struct {
	model.Params
	// 系统自动生成
	_queryCpRequest *Querycprequest
}

// NewAlibabaAscpUopTaobaoWaybillQuerycpRequest 初始化AlibabaAscpUopTaobaoWaybillQuerycpAPIRequest对象
func NewAlibabaAscpUopTaobaoWaybillQuerycpRequest() *AlibabaAscpUopTaobaoWaybillQuerycpAPIRequest {
	return &AlibabaAscpUopTaobaoWaybillQuerycpAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopTaobaoWaybillQuerycpAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.taobao.waybill.querycp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopTaobaoWaybillQuerycpAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is QueryCpRequest Setter
// 系统自动生成
func (r *AlibabaAscpUopTaobaoWaybillQuerycpAPIRequest) SetQueryCpRequest(_queryCpRequest *Querycprequest) error {
	r._queryCpRequest = _queryCpRequest
	r.Set("query_cp_request", _queryCpRequest)
	return nil
}

// Get QueryCpRequest Getter
func (r AlibabaAscpUopTaobaoWaybillQuerycpAPIRequest) GetQueryCpRequest() *Querycprequest {
	return r._queryCpRequest
}
