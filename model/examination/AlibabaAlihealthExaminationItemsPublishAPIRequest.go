package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationItemsPublishAPIRequest 单项/加项包信息同步 API请求
// alibaba.alihealth.examination.items.publish
//
// 体检机构对接_单项/加项包信息发布／更新
type AlibabaAlihealthExaminationItemsPublishAPIRequest struct {
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

// NewAlibabaAlihealthExaminationItemsPublishRequest 初始化AlibabaAlihealthExaminationItemsPublishAPIRequest对象
func NewAlibabaAlihealthExaminationItemsPublishRequest() *AlibabaAlihealthExaminationItemsPublishAPIRequest {
	return &AlibabaAlihealthExaminationItemsPublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.items.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetIsvPackages is IsvPackages Setter
// 套餐列表
func (r *AlibabaAlihealthExaminationItemsPublishAPIRequest) SetIsvPackages(_isvPackages []IsvPackage) error {
	r._isvPackages = _isvPackages
	r.Set("isv_packages", _isvPackages)
	return nil
}

// GetIsvPackages IsvPackages Getter
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetIsvPackages() []IsvPackage {
	return r._isvPackages
}

// SetIsvPackRelationDTOS is IsvPackRelationDTOS Setter
// 加项包关系列表
func (r *AlibabaAlihealthExaminationItemsPublishAPIRequest) SetIsvPackRelationDTOS(_isvPackRelationDTOS []IsvPackRelationDto) error {
	r._isvPackRelationDTOS = _isvPackRelationDTOS
	r.Set("isv_pack_relation_d_t_o_s", _isvPackRelationDTOS)
	return nil
}

// GetIsvPackRelationDTOS IsvPackRelationDTOS Getter
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetIsvPackRelationDTOS() []IsvPackRelationDto {
	return r._isvPackRelationDTOS
}

// SetIsvItemRelationDTOS is IsvItemRelationDTOS Setter
// 单项之间关系
func (r *AlibabaAlihealthExaminationItemsPublishAPIRequest) SetIsvItemRelationDTOS(_isvItemRelationDTOS []IsvItemRelationDto) error {
	r._isvItemRelationDTOS = _isvItemRelationDTOS
	r.Set("isv_item_relation_d_t_o_s", _isvItemRelationDTOS)
	return nil
}

// GetIsvItemRelationDTOS IsvItemRelationDTOS Getter
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetIsvItemRelationDTOS() []IsvItemRelationDto {
	return r._isvItemRelationDTOS
}

// SetHospitalCodes is HospitalCodes Setter
// 体检机构标识
func (r *AlibabaAlihealthExaminationItemsPublishAPIRequest) SetHospitalCodes(_hospitalCodes []string) error {
	r._hospitalCodes = _hospitalCodes
	r.Set("hospital_codes", _hospitalCodes)
	return nil
}

// GetHospitalCodes HospitalCodes Getter
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetHospitalCodes() []string {
	return r._hospitalCodes
}

// SetIsvItemPackDTOS is IsvItemPackDTOS Setter
// 加项包列表
func (r *AlibabaAlihealthExaminationItemsPublishAPIRequest) SetIsvItemPackDTOS(_isvItemPackDTOS []IsvItemPackDto) error {
	r._isvItemPackDTOS = _isvItemPackDTOS
	r.Set("isv_item_pack_d_t_o_s", _isvItemPackDTOS)
	return nil
}

// GetIsvItemPackDTOS IsvItemPackDTOS Getter
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetIsvItemPackDTOS() []IsvItemPackDto {
	return r._isvItemPackDTOS
}

// SetIsvItemDTOS is IsvItemDTOS Setter
// 单项信息列表
func (r *AlibabaAlihealthExaminationItemsPublishAPIRequest) SetIsvItemDTOS(_isvItemDTOS []IsvItemDto) error {
	r._isvItemDTOS = _isvItemDTOS
	r.Set("isv_item_d_t_o_s", _isvItemDTOS)
	return nil
}

// GetIsvItemDTOS IsvItemDTOS Getter
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetIsvItemDTOS() []IsvItemDto {
	return r._isvItemDTOS
}

// SetGroupId is GroupId Setter
// 商品id，机构保证全局唯一
func (r *AlibabaAlihealthExaminationItemsPublishAPIRequest) SetGroupId(_groupId string) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetGroupId() string {
	return r._groupId
}
