package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaokoubeitribeopenverifycodeapplyAPIRequest 口碑综合体手机号获取验证码 API请求
// taobao.koubei.tribe.open.verify.code.apply
//
// 口碑综合体通过手机号获取验证码对外开放接口
type TaobaokoubeitribeopenverifycodeapplyAPIRequest struct {
	model.Params
	// 数据集id
	_dataSetId string
	// 手机号
	_phone string
}

// NewTaobaokoubeitribeopenverifycodeapplyRequest 初始化TaobaokoubeitribeopenverifycodeapplyAPIRequest对象
func NewTaobaokoubeitribeopenverifycodeapplyRequest() *TaobaokoubeitribeopenverifycodeapplyAPIRequest {
	return &TaobaokoubeitribeopenverifycodeapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaokoubeitribeopenverifycodeapplyAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.tribe.open.verify.code.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaokoubeitribeopenverifycodeapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaokoubeitribeopenverifycodeapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDataSetId is DataSetId Setter
// 数据集id
func (r *TaobaokoubeitribeopenverifycodeapplyAPIRequest) SetDataSetId(_dataSetId string) error {
	r._dataSetId = _dataSetId
	r.Set("data_set_id", _dataSetId)
	return nil
}

// GetDataSetId DataSetId Getter
func (r TaobaokoubeitribeopenverifycodeapplyAPIRequest) GetDataSetId() string {
	return r._dataSetId
}

// SetPhone is Phone Setter
// 手机号
func (r *TaobaokoubeitribeopenverifycodeapplyAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r TaobaokoubeitribeopenverifycodeapplyAPIRequest) GetPhone() string {
	return r._phone
}
