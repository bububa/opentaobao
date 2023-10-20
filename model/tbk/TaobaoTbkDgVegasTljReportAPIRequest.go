package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgvegastljreportAPIRequest 淘宝客-推广者-淘礼金效果数据 API请求
// taobao.tbk.dg.vegas.tlj.report
//
// 淘宝客推广者可查询淘礼金发放和使用等效果数据，只提供近150天的数据
type TaobaotbkdgvegastljreportAPIRequest struct {
	model.Params
	// 创建淘礼金时返回的rightsId
	_rightsId string
	// adzoneId
	_adzoneId int64
}

// NewTaobaotbkdgvegastljreportRequest 初始化TaobaotbkdgvegastljreportAPIRequest对象
func NewTaobaotbkdgvegastljreportRequest() *TaobaotbkdgvegastljreportAPIRequest {
	return &TaobaotbkdgvegastljreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkdgvegastljreportAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.vegas.tlj.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkdgvegastljreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkdgvegastljreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRightsId is RightsId Setter
// 创建淘礼金时返回的rightsId
func (r *TaobaotbkdgvegastljreportAPIRequest) SetRightsId(_rightsId string) error {
	r._rightsId = _rightsId
	r.Set("rights_id", _rightsId)
	return nil
}

// GetRightsId RightsId Getter
func (r TaobaotbkdgvegastljreportAPIRequest) GetRightsId() string {
	return r._rightsId
}

// SetAdzoneId is AdzoneId Setter
// adzoneId
func (r *TaobaotbkdgvegastljreportAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaotbkdgvegastljreportAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}
