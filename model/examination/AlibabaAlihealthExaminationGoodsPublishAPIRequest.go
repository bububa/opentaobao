package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationgoodspublishAPIRequest 体检机构对接_商品发布／更新 API请求
// alibaba.alihealth.examination.goods.publish
//
// 体检机构对接_商品发布／更新
type AlibabaalihealthexaminationgoodspublishAPIRequest struct {
	model.Params
	// 套餐列表
	_packageList []Package
	// 商品id，机构保证全局唯一
	_groupId string
	// 商品名称
	_groupName string
	// 操作类型: publish=发布，update=更新, 更新时如果套餐列表内套餐有移除则视为删除套餐
	_type string
	// 最多200个字，界面对应商品详情页描述
	_goodsDesc string
	// 最多256个字，界面对应列表文字；为空是自动取套餐的package_detail字段值
	_targetGroup string
	// 标签值给C，传错C端将无法透出
	_label string
	// 商品类目，1：体检 ，2：核酸，4 ：健康证
	_categoryId string
	// 0自营商品，1平台商品
	_mode string
	// 类目ID，填入叶子类目ID，儿童体检: 20210204000004, 中青年体检: 20210204000005, 老年体检: 20210204000006, 证件体检（含入职）: 20210204000007, 核酸检测（到店服务）: 20210204000008, 专科服务（不包含核酸检测）: 20210204000009, 上门检测: 202102040000010, 上门护理: 202102040000011, 上门体检 202102040000012
	_backendCategoryId int64
	// 预约服务版本号, 控制预约链路, 仅支持新增, 不支持更新; 1: 路由至预约系统，0:系统默认；
	_reservationApiVersion int64
}

// NewAlibabaalihealthexaminationgoodspublishRequest 初始化AlibabaalihealthexaminationgoodspublishAPIRequest对象
func NewAlibabaalihealthexaminationgoodspublishRequest() *AlibabaalihealthexaminationgoodspublishAPIRequest {
	return &AlibabaalihealthexaminationgoodspublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationgoodspublishAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.goods.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationgoodspublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationgoodspublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPackageList is PackageList Setter
// 套餐列表
func (r *AlibabaalihealthexaminationgoodspublishAPIRequest) SetPackageList(_packageList []Package) error {
	r._packageList = _packageList
	r.Set("package_list", _packageList)
	return nil
}

// GetPackageList PackageList Getter
func (r AlibabaalihealthexaminationgoodspublishAPIRequest) GetPackageList() []Package {
	return r._packageList
}

// SetGroupId is GroupId Setter
// 商品id，机构保证全局唯一
func (r *AlibabaalihealthexaminationgoodspublishAPIRequest) SetGroupId(_groupId string) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabaalihealthexaminationgoodspublishAPIRequest) GetGroupId() string {
	return r._groupId
}

// SetGroupName is GroupName Setter
// 商品名称
func (r *AlibabaalihealthexaminationgoodspublishAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r AlibabaalihealthexaminationgoodspublishAPIRequest) GetGroupName() string {
	return r._groupName
}

// SetType is Type Setter
// 操作类型: publish=发布，update=更新, 更新时如果套餐列表内套餐有移除则视为删除套餐
func (r *AlibabaalihealthexaminationgoodspublishAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaalihealthexaminationgoodspublishAPIRequest) GetType() string {
	return r._type
}

// SetGoodsDesc is GoodsDesc Setter
// 最多200个字，界面对应商品详情页描述
func (r *AlibabaalihealthexaminationgoodspublishAPIRequest) SetGoodsDesc(_goodsDesc string) error {
	r._goodsDesc = _goodsDesc
	r.Set("goods_desc", _goodsDesc)
	return nil
}

// GetGoodsDesc GoodsDesc Getter
func (r AlibabaalihealthexaminationgoodspublishAPIRequest) GetGoodsDesc() string {
	return r._goodsDesc
}

// SetTargetGroup is TargetGroup Setter
// 最多256个字，界面对应列表文字；为空是自动取套餐的package_detail字段值
func (r *AlibabaalihealthexaminationgoodspublishAPIRequest) SetTargetGroup(_targetGroup string) error {
	r._targetGroup = _targetGroup
	r.Set("target_group", _targetGroup)
	return nil
}

// GetTargetGroup TargetGroup Getter
func (r AlibabaalihealthexaminationgoodspublishAPIRequest) GetTargetGroup() string {
	return r._targetGroup
}

// SetLabel is Label Setter
// 标签值给C，传错C端将无法透出
func (r *AlibabaalihealthexaminationgoodspublishAPIRequest) SetLabel(_label string) error {
	r._label = _label
	r.Set("label", _label)
	return nil
}

// GetLabel Label Getter
func (r AlibabaalihealthexaminationgoodspublishAPIRequest) GetLabel() string {
	return r._label
}

// SetCategoryId is CategoryId Setter
// 商品类目，1：体检 ，2：核酸，4 ：健康证
func (r *AlibabaalihealthexaminationgoodspublishAPIRequest) SetCategoryId(_categoryId string) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r AlibabaalihealthexaminationgoodspublishAPIRequest) GetCategoryId() string {
	return r._categoryId
}

// SetMode is Mode Setter
// 0自营商品，1平台商品
func (r *AlibabaalihealthexaminationgoodspublishAPIRequest) SetMode(_mode string) error {
	r._mode = _mode
	r.Set("mode", _mode)
	return nil
}

// GetMode Mode Getter
func (r AlibabaalihealthexaminationgoodspublishAPIRequest) GetMode() string {
	return r._mode
}

// SetBackendCategoryId is BackendCategoryId Setter
// 类目ID，填入叶子类目ID，儿童体检: 20210204000004, 中青年体检: 20210204000005, 老年体检: 20210204000006, 证件体检（含入职）: 20210204000007, 核酸检测（到店服务）: 20210204000008, 专科服务（不包含核酸检测）: 20210204000009, 上门检测: 202102040000010, 上门护理: 202102040000011, 上门体检 202102040000012
func (r *AlibabaalihealthexaminationgoodspublishAPIRequest) SetBackendCategoryId(_backendCategoryId int64) error {
	r._backendCategoryId = _backendCategoryId
	r.Set("backend_category_id", _backendCategoryId)
	return nil
}

// GetBackendCategoryId BackendCategoryId Getter
func (r AlibabaalihealthexaminationgoodspublishAPIRequest) GetBackendCategoryId() int64 {
	return r._backendCategoryId
}

// SetReservationApiVersion is ReservationApiVersion Setter
// 预约服务版本号, 控制预约链路, 仅支持新增, 不支持更新; 1: 路由至预约系统，0:系统默认；
func (r *AlibabaalihealthexaminationgoodspublishAPIRequest) SetReservationApiVersion(_reservationApiVersion int64) error {
	r._reservationApiVersion = _reservationApiVersion
	r.Set("reservation_api_version", _reservationApiVersion)
	return nil
}

// GetReservationApiVersion ReservationApiVersion Getter
func (r AlibabaalihealthexaminationgoodspublishAPIRequest) GetReservationApiVersion() int64 {
	return r._reservationApiVersion
}
