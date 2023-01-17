package ticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTicketScenicBindAPIRequest 【门票API2.0】门票景点绑定接口 API请求
// alitrip.ticket.scenic.bind
//
// 门票景点绑定接口，用于建立阿里标准景点id与商家系统景点id的映射关系。该接口同时支持新建和修改映射关系，当用户没有为ali_scenic_id建立过映射关系时，则判断为新建映射关系，否则为修改。可以通过设置update_out_scenic_id来修改ali_scenic_id与out_scenic_id的映射关系。
type AlitripTicketScenicBindAPIRequest struct {
	model.Params
	// 商户景点城市
	_city string
	// 商户景点地址
	_address string
	// 商户景点名称
	_outScenicName string
	// 商户景点省份
	_province string
	// 必填，商户系统中的景点编码，用于与ali_scenic_id建立映射关系
	_outScenicId string
	// 可选，如果需要更新out_scenic_id与ali_scenic_id的映射关系时 需要填写
	_updateOutScenicId string
	// 必填，阿里旅行对应的景点编码
	_aliScenicId int64
}

// NewAlitripTicketScenicBindRequest 初始化AlitripTicketScenicBindAPIRequest对象
func NewAlitripTicketScenicBindRequest() *AlitripTicketScenicBindAPIRequest {
	return &AlitripTicketScenicBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTicketScenicBindAPIRequest) GetApiMethodName() string {
	return "alitrip.ticket.scenic.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTicketScenicBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTicketScenicBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCity is City Setter
// 商户景点城市
func (r *AlitripTicketScenicBindAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r AlitripTicketScenicBindAPIRequest) GetCity() string {
	return r._city
}

// SetAddress is Address Setter
// 商户景点地址
func (r *AlitripTicketScenicBindAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r AlitripTicketScenicBindAPIRequest) GetAddress() string {
	return r._address
}

// SetOutScenicName is OutScenicName Setter
// 商户景点名称
func (r *AlitripTicketScenicBindAPIRequest) SetOutScenicName(_outScenicName string) error {
	r._outScenicName = _outScenicName
	r.Set("out_scenic_name", _outScenicName)
	return nil
}

// GetOutScenicName OutScenicName Getter
func (r AlitripTicketScenicBindAPIRequest) GetOutScenicName() string {
	return r._outScenicName
}

// SetProvince is Province Setter
// 商户景点省份
func (r *AlitripTicketScenicBindAPIRequest) SetProvince(_province string) error {
	r._province = _province
	r.Set("province", _province)
	return nil
}

// GetProvince Province Getter
func (r AlitripTicketScenicBindAPIRequest) GetProvince() string {
	return r._province
}

// SetOutScenicId is OutScenicId Setter
// 必填，商户系统中的景点编码，用于与ali_scenic_id建立映射关系
func (r *AlitripTicketScenicBindAPIRequest) SetOutScenicId(_outScenicId string) error {
	r._outScenicId = _outScenicId
	r.Set("out_scenic_id", _outScenicId)
	return nil
}

// GetOutScenicId OutScenicId Getter
func (r AlitripTicketScenicBindAPIRequest) GetOutScenicId() string {
	return r._outScenicId
}

// SetUpdateOutScenicId is UpdateOutScenicId Setter
// 可选，如果需要更新out_scenic_id与ali_scenic_id的映射关系时 需要填写
func (r *AlitripTicketScenicBindAPIRequest) SetUpdateOutScenicId(_updateOutScenicId string) error {
	r._updateOutScenicId = _updateOutScenicId
	r.Set("update_out_scenic_id", _updateOutScenicId)
	return nil
}

// GetUpdateOutScenicId UpdateOutScenicId Getter
func (r AlitripTicketScenicBindAPIRequest) GetUpdateOutScenicId() string {
	return r._updateOutScenicId
}

// SetAliScenicId is AliScenicId Setter
// 必填，阿里旅行对应的景点编码
func (r *AlitripTicketScenicBindAPIRequest) SetAliScenicId(_aliScenicId int64) error {
	r._aliScenicId = _aliScenicId
	r.Set("ali_scenic_id", _aliScenicId)
	return nil
}

// GetAliScenicId AliScenicId Getter
func (r AlitripTicketScenicBindAPIRequest) GetAliScenicId() int64 {
	return r._aliScenicId
}
