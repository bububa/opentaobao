package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationitemspublishAPIRequest 单项/加项包信息同步 API请求
// alibaba.alihealth.examination.items.publish
//
// 体检机构对接_单项/加项包信息发布／更新
type AlibabaalihealthexaminationitemspublishAPIRequest struct {
	model.Params
	// 套餐列表
	_isvPackages []IsvPackage
	// 加项包关系列表
	_isvPackRelationDTOS []IsvPackRelationDto
	// 单项之间关系
	_isvItemRelationDTOS []IsvItemRelationDto
	// 体检机构标识
	_hospitalCodes []string
	// 加项包列表
	_isvItemPackDTOS []IsvItemPackDto
	// 单项信息列表
	_isvItemDTOS []IsvItemDto
	// 商品id，机构保证全局唯一
	_groupId string
}

// NewAlibabaalihealthexaminationitemspublishRequest 初始化AlibabaalihealthexaminationitemspublishAPIRequest对象
func NewAlibabaalihealthexaminationitemspublishRequest() *AlibabaalihealthexaminationitemspublishAPIRequest {
	return &AlibabaalihealthexaminationitemspublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationitemspublishAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.items.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationitemspublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationitemspublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvPackages is IsvPackages Setter
// 套餐列表
func (r *AlibabaalihealthexaminationitemspublishAPIRequest) SetIsvPackages(_isvPackages []IsvPackage) error {
	r._isvPackages = _isvPackages
	r.Set("isv_packages", _isvPackages)
	return nil
}

// GetIsvPackages IsvPackages Getter
func (r AlibabaalihealthexaminationitemspublishAPIRequest) GetIsvPackages() []IsvPackage {
	return r._isvPackages
}

// SetIsvPackRelationDTOS is IsvPackRelationDTOS Setter
// 加项包关系列表
func (r *AlibabaalihealthexaminationitemspublishAPIRequest) SetIsvPackRelationDTOS(_isvPackRelationDTOS []IsvPackRelationDto) error {
	r._isvPackRelationDTOS = _isvPackRelationDTOS
	r.Set("isv_pack_relation_d_t_o_s", _isvPackRelationDTOS)
	return nil
}

// GetIsvPackRelationDTOS IsvPackRelationDTOS Getter
func (r AlibabaalihealthexaminationitemspublishAPIRequest) GetIsvPackRelationDTOS() []IsvPackRelationDto {
	return r._isvPackRelationDTOS
}

// SetIsvItemRelationDTOS is IsvItemRelationDTOS Setter
// 单项之间关系
func (r *AlibabaalihealthexaminationitemspublishAPIRequest) SetIsvItemRelationDTOS(_isvItemRelationDTOS []IsvItemRelationDto) error {
	r._isvItemRelationDTOS = _isvItemRelationDTOS
	r.Set("isv_item_relation_d_t_o_s", _isvItemRelationDTOS)
	return nil
}

// GetIsvItemRelationDTOS IsvItemRelationDTOS Getter
func (r AlibabaalihealthexaminationitemspublishAPIRequest) GetIsvItemRelationDTOS() []IsvItemRelationDto {
	return r._isvItemRelationDTOS
}

// SetHospitalCodes is HospitalCodes Setter
// 体检机构标识
func (r *AlibabaalihealthexaminationitemspublishAPIRequest) SetHospitalCodes(_hospitalCodes []string) error {
	r._hospitalCodes = _hospitalCodes
	r.Set("hospital_codes", _hospitalCodes)
	return nil
}

// GetHospitalCodes HospitalCodes Getter
func (r AlibabaalihealthexaminationitemspublishAPIRequest) GetHospitalCodes() []string {
	return r._hospitalCodes
}

// SetIsvItemPackDTOS is IsvItemPackDTOS Setter
// 加项包列表
func (r *AlibabaalihealthexaminationitemspublishAPIRequest) SetIsvItemPackDTOS(_isvItemPackDTOS []IsvItemPackDto) error {
	r._isvItemPackDTOS = _isvItemPackDTOS
	r.Set("isv_item_pack_d_t_o_s", _isvItemPackDTOS)
	return nil
}

// GetIsvItemPackDTOS IsvItemPackDTOS Getter
func (r AlibabaalihealthexaminationitemspublishAPIRequest) GetIsvItemPackDTOS() []IsvItemPackDto {
	return r._isvItemPackDTOS
}

// SetIsvItemDTOS is IsvItemDTOS Setter
// 单项信息列表
func (r *AlibabaalihealthexaminationitemspublishAPIRequest) SetIsvItemDTOS(_isvItemDTOS []IsvItemDto) error {
	r._isvItemDTOS = _isvItemDTOS
	r.Set("isv_item_d_t_o_s", _isvItemDTOS)
	return nil
}

// GetIsvItemDTOS IsvItemDTOS Getter
func (r AlibabaalihealthexaminationitemspublishAPIRequest) GetIsvItemDTOS() []IsvItemDto {
	return r._isvItemDTOS
}

// SetGroupId is GroupId Setter
// 商品id，机构保证全局唯一
func (r *AlibabaalihealthexaminationitemspublishAPIRequest) SetGroupId(_groupId string) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabaalihealthexaminationitemspublishAPIRequest) GetGroupId() string {
	return r._groupId
}
