package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationGoodsPublishAPIRequest
体检机构对接_商品发布／更新 API请求
alibaba.alihealth.examination.goods.publish

体检机构对接_商品发布／更新 */
type AlibabaAlihealthExaminationGoodsPublishAPIRequest struct {
	model.Params
	// 商品id，机构保证全局唯一
	_groupId string
	// 商品名称
	_groupName string
	// 套餐列表
	_packageList []Package
	// 操作类型: publish=发布，update=更新
	_type string
	// 最多200个字，界面对应商品详情页描述
	_goodsDesc string
	// 最多256个字，界面对应列表文字
	_targetGroup string
	// 联调中正式上线前标签给B；联调后正式上线后标签给C
	_label string
	// 商品类目，1：体检 ，2：核酸，4 ：健康证
	_categoryId string
	// 0自营商品，1平台商品
	_mode string
	// 类目ID，填入叶子类目ID，儿童体检: 20210204000004, 中青年体检: 20210204000005, 老年体检: 20210204000006, 证件体检（含入职）: 20210204000007, 核酸检测（到店服务）: 20210204000008, 专科服务（不包含核酸检测）: 20210204000009, 上门检测: 202102040000010, 上门护理: 202102040000011, 上门体检 202102040000012
	_backendCategoryId int64
}

// NewAlibabaAlihealthExaminationGoodsPublishRequest 初始化AlibabaAlihealthExaminationGoodsPublishAPIRequest对象
func NewAlibabaAlihealthExaminationGoodsPublishRequest() *AlibabaAlihealthExaminationGoodsPublishAPIRequest {
	return &AlibabaAlihealthExaminationGoodsPublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationGoodsPublishAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.goods.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationGoodsPublishAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is GroupId Setter
// 商品id，机构保证全局唯一
func (r *AlibabaAlihealthExaminationGoodsPublishAPIRequest) SetGroupId(_groupId string) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// Get GroupId Getter
func (r AlibabaAlihealthExaminationGoodsPublishAPIRequest) GetGroupId() string {
	return r._groupId
}

// Set is GroupName Setter
// 商品名称
func (r *AlibabaAlihealthExaminationGoodsPublishAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// Get GroupName Getter
func (r AlibabaAlihealthExaminationGoodsPublishAPIRequest) GetGroupName() string {
	return r._groupName
}

// Set is PackageList Setter
// 套餐列表
func (r *AlibabaAlihealthExaminationGoodsPublishAPIRequest) SetPackageList(_packageList []Package) error {
	r._packageList = _packageList
	r.Set("package_list", _packageList)
	return nil
}

// Get PackageList Getter
func (r AlibabaAlihealthExaminationGoodsPublishAPIRequest) GetPackageList() []Package {
	return r._packageList
}

// Set is Type Setter
// 操作类型: publish=发布，update=更新
func (r *AlibabaAlihealthExaminationGoodsPublishAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r AlibabaAlihealthExaminationGoodsPublishAPIRequest) GetType() string {
	return r._type
}

// Set is GoodsDesc Setter
// 最多200个字，界面对应商品详情页描述
func (r *AlibabaAlihealthExaminationGoodsPublishAPIRequest) SetGoodsDesc(_goodsDesc string) error {
	r._goodsDesc = _goodsDesc
	r.Set("goods_desc", _goodsDesc)
	return nil
}

// Get GoodsDesc Getter
func (r AlibabaAlihealthExaminationGoodsPublishAPIRequest) GetGoodsDesc() string {
	return r._goodsDesc
}

// Set is TargetGroup Setter
// 最多256个字，界面对应列表文字
func (r *AlibabaAlihealthExaminationGoodsPublishAPIRequest) SetTargetGroup(_targetGroup string) error {
	r._targetGroup = _targetGroup
	r.Set("target_group", _targetGroup)
	return nil
}

// Get TargetGroup Getter
func (r AlibabaAlihealthExaminationGoodsPublishAPIRequest) GetTargetGroup() string {
	return r._targetGroup
}

// Set is Label Setter
// 联调中正式上线前标签给B；联调后正式上线后标签给C
func (r *AlibabaAlihealthExaminationGoodsPublishAPIRequest) SetLabel(_label string) error {
	r._label = _label
	r.Set("label", _label)
	return nil
}

// Get Label Getter
func (r AlibabaAlihealthExaminationGoodsPublishAPIRequest) GetLabel() string {
	return r._label
}

// Set is CategoryId Setter
// 商品类目，1：体检 ，2：核酸，4 ：健康证
func (r *AlibabaAlihealthExaminationGoodsPublishAPIRequest) SetCategoryId(_categoryId string) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// Get CategoryId Getter
func (r AlibabaAlihealthExaminationGoodsPublishAPIRequest) GetCategoryId() string {
	return r._categoryId
}

// Set is Mode Setter
// 0自营商品，1平台商品
func (r *AlibabaAlihealthExaminationGoodsPublishAPIRequest) SetMode(_mode string) error {
	r._mode = _mode
	r.Set("mode", _mode)
	return nil
}

// Get Mode Getter
func (r AlibabaAlihealthExaminationGoodsPublishAPIRequest) GetMode() string {
	return r._mode
}

// Set is BackendCategoryId Setter
// 类目ID，填入叶子类目ID，儿童体检: 20210204000004, 中青年体检: 20210204000005, 老年体检: 20210204000006, 证件体检（含入职）: 20210204000007, 核酸检测（到店服务）: 20210204000008, 专科服务（不包含核酸检测）: 20210204000009, 上门检测: 202102040000010, 上门护理: 202102040000011, 上门体检 202102040000012
func (r *AlibabaAlihealthExaminationGoodsPublishAPIRequest) SetBackendCategoryId(_backendCategoryId int64) error {
	r._backendCategoryId = _backendCategoryId
	r.Set("backend_category_id", _backendCategoryId)
	return nil
}

// Get BackendCategoryId Getter
func (r AlibabaAlihealthExaminationGoodsPublishAPIRequest) GetBackendCategoryId() int64 {
	return r._backendCategoryId
}
