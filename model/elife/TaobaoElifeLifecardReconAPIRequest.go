package elife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoelifelifecardreconAPIRequest 查询对账文件地址接口 API请求
// taobao.elife.lifecard.recon
//
// 查询对账文件地址接口
type TaobaoelifelifecardreconAPIRequest struct {
	model.Params
	// 对账日期(YYYYMMDD)
	_opDate string
}

// NewTaobaoelifelifecardreconRequest 初始化TaobaoelifelifecardreconAPIRequest对象
func NewTaobaoelifelifecardreconRequest() *TaobaoelifelifecardreconAPIRequest {
	return &TaobaoelifelifecardreconAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoelifelifecardreconAPIRequest) GetApiMethodName() string {
	return "taobao.elife.lifecard.recon"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoelifelifecardreconAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoelifelifecardreconAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpDate is OpDate Setter
// 对账日期(YYYYMMDD)
func (r *TaobaoelifelifecardreconAPIRequest) SetOpDate(_opDate string) error {
	r._opDate = _opDate
	r.Set("op_date", _opDate)
	return nil
}

// GetOpDate OpDate Getter
func (r TaobaoelifelifecardreconAPIRequest) GetOpDate() string {
	return r._opDate
}
