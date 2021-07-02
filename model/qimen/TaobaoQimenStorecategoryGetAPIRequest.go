package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStorecategoryGetAPIRequest 门店类目获取接口 API请求
// taobao.qimen.storecategory.get
//
// 商家在ERP中调用该接口，获取门店类目
type TaobaoQimenStorecategoryGetAPIRequest struct {
	model.Params
	// 备注
	_remark string
}

// NewTaobaoQimenStorecategoryGetRequest 初始化TaobaoQimenStorecategoryGetAPIRequest对象
func NewTaobaoQimenStorecategoryGetRequest() *TaobaoQimenStorecategoryGetAPIRequest {
	return &TaobaoQimenStorecategoryGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenStorecategoryGetAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.storecategory.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenStorecategoryGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRemark is Remark Setter
// 备注
func (r *TaobaoQimenStorecategoryGetAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoQimenStorecategoryGetAPIRequest) GetRemark() string {
	return r._remark
}
