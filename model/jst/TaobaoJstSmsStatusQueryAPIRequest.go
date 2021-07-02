package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsStatusQueryAPIRequest 聚石塔公众号状态查询 API请求
// taobao.jst.sms.status.query
//
// 聚石塔公众号状态查询
type TaobaoJstSmsStatusQueryAPIRequest struct {
	model.Params
	// 公众号状态信息查询请求
	_officialAccountStatusQueryRequest *JstBaseRequest
}

// NewTaobaoJstSmsStatusQueryRequest 初始化TaobaoJstSmsStatusQueryAPIRequest对象
func NewTaobaoJstSmsStatusQueryRequest() *TaobaoJstSmsStatusQueryAPIRequest {
	return &TaobaoJstSmsStatusQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsStatusQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.status.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsStatusQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OfficialAccountStatusQueryRequest Setter
// 公众号状态信息查询请求
func (r *TaobaoJstSmsStatusQueryAPIRequest) SetOfficialAccountStatusQueryRequest(_officialAccountStatusQueryRequest *JstBaseRequest) error {
	r._officialAccountStatusQueryRequest = _officialAccountStatusQueryRequest
	r.Set("official_account_status_query_request", _officialAccountStatusQueryRequest)
	return nil
}

// Get OfficialAccountStatusQueryRequest Getter
func (r TaobaoJstSmsStatusQueryAPIRequest) GetOfficialAccountStatusQueryRequest() *JstBaseRequest {
	return r._officialAccountStatusQueryRequest
}
