package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaonewretaildivisionrecordlistgetAPIRequest 导购分佣明细列表 API请求
// taobao.newretail.division.record.list.get
//
// 提供分页查询导购分佣明细的能力
type TaobaonewretaildivisionrecordlistgetAPIRequest struct {
	model.Params
	// 入参
	_param *TopDivisionRecordReqDto
}

// NewTaobaonewretaildivisionrecordlistgetRequest 初始化TaobaonewretaildivisionrecordlistgetAPIRequest对象
func NewTaobaonewretaildivisionrecordlistgetRequest() *TaobaonewretaildivisionrecordlistgetAPIRequest {
	return &TaobaonewretaildivisionrecordlistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaonewretaildivisionrecordlistgetAPIRequest) GetApiMethodName() string {
	return "taobao.newretail.division.record.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaonewretaildivisionrecordlistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaonewretaildivisionrecordlistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *TaobaonewretaildivisionrecordlistgetAPIRequest) SetParam(_param *TopDivisionRecordReqDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaonewretaildivisionrecordlistgetAPIRequest) GetParam() *TopDivisionRecordReqDto {
	return r._param
}
