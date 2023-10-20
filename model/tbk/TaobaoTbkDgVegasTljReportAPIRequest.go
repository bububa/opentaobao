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
	_rightsid string
	// adzoneId
	_adzoneid int64
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

// SetRightsid is Rightsid Setter
// 创建淘礼金时返回的rightsId
func (r *TaobaotbkdgvegastljreportAPIRequest) SetRightsid(_rightsid string) error {
	r._rightsid = _rightsid
	r.Set("rights_id", _rightsid)
	return nil
}

// GetRightsid Rightsid Getter
func (r TaobaotbkdgvegastljreportAPIRequest) GetRightsid() string {
	return r._rightsid
}

// SetAdzoneid is Adzoneid Setter
// adzoneId
func (r *TaobaotbkdgvegastljreportAPIRequest) SetAdzoneid(_adzoneid int64) error {
	r._adzoneid = _adzoneid
	r.Set("adzone_id", _adzoneid)
	return nil
}

// GetAdzoneid Adzoneid Getter
func (r TaobaotbkdgvegastljreportAPIRequest) GetAdzoneid() int64 {
	return r._adzoneid
}
