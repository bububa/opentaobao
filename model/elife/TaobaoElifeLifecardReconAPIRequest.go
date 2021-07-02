package elife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoElifeLifecardReconAPIRequest 查询对账文件地址接口 API请求
// taobao.elife.lifecard.recon
//
// 查询对账文件地址接口
type TaobaoElifeLifecardReconAPIRequest struct {
	model.Params
	// 对账日期(YYYYMMDD)
	_opDate string
}

// NewTaobaoElifeLifecardReconRequest 初始化TaobaoElifeLifecardReconAPIRequest对象
func NewTaobaoElifeLifecardReconRequest() *TaobaoElifeLifecardReconAPIRequest {
	return &TaobaoElifeLifecardReconAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoElifeLifecardReconAPIRequest) GetApiMethodName() string {
	return "taobao.elife.lifecard.recon"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoElifeLifecardReconAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OpDate Setter
// 对账日期(YYYYMMDD)
func (r *TaobaoElifeLifecardReconAPIRequest) SetOpDate(_opDate string) error {
	r._opDate = _opDate
	r.Set("op_date", _opDate)
	return nil
}

// Get OpDate Getter
func (r TaobaoElifeLifecardReconAPIRequest) GetOpDate() string {
	return r._opDate
}
