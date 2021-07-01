package yunosappstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosAppstoreOpenGetadsAPIRequest
获取外投广告 API请求
yunos.appstore.open.getads

将广告外投给外部合作伙伴 */
type YunosAppstoreOpenGetadsAPIRequest struct {
	model.Params
	// 请求id
	_rid string
	// 指定广告分类
	_cats []string
	// 是否排除已安装
	_excludeInstall bool
	// 场景或页面标识
	_caseId string
	// ssp标识
	_ssp string
	// 结算类型
	_feeType string
	// 客户端来源ip
	_clientIp string
	// 广告指定包名
	_pkgs []string
	// 客户端版本号
	_clientVerCode int64
	// 是否映射到uuid
	_tryMapToUuid bool
	// 排除包名列表
	_excludePkgs []string
	// 设备唯一标识
	_deviceId string
	// 广告数量
	_size int64
	// 排除分类
	_excludeCats []string
	// 创意模板id列表
	_templateIds []int64
	// 广告底价
	_mrp int64
	// 请求特征集
	_options int64
}

// NewYunosAppstoreOpenGetadsRequest 初始化YunosAppstoreOpenGetadsAPIRequest对象
func NewYunosAppstoreOpenGetadsRequest() *YunosAppstoreOpenGetadsAPIRequest {
	return &YunosAppstoreOpenGetadsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosAppstoreOpenGetadsAPIRequest) GetApiMethodName() string {
	return "yunos.appstore.open.getads"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosAppstoreOpenGetadsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rid Setter
// 请求id
func (r *YunosAppstoreOpenGetadsAPIRequest) SetRid(_rid string) error {
	r._rid = _rid
	r.Set("rid", _rid)
	return nil
}

// Get Rid Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetRid() string {
	return r._rid
}

// Set is Cats Setter
// 指定广告分类
func (r *YunosAppstoreOpenGetadsAPIRequest) SetCats(_cats []string) error {
	r._cats = _cats
	r.Set("cats", _cats)
	return nil
}

// Get Cats Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetCats() []string {
	return r._cats
}

// Set is ExcludeInstall Setter
// 是否排除已安装
func (r *YunosAppstoreOpenGetadsAPIRequest) SetExcludeInstall(_excludeInstall bool) error {
	r._excludeInstall = _excludeInstall
	r.Set("exclude_install", _excludeInstall)
	return nil
}

// Get ExcludeInstall Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetExcludeInstall() bool {
	return r._excludeInstall
}

// Set is CaseId Setter
// 场景或页面标识
func (r *YunosAppstoreOpenGetadsAPIRequest) SetCaseId(_caseId string) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// Get CaseId Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetCaseId() string {
	return r._caseId
}

// Set is Ssp Setter
// ssp标识
func (r *YunosAppstoreOpenGetadsAPIRequest) SetSsp(_ssp string) error {
	r._ssp = _ssp
	r.Set("ssp", _ssp)
	return nil
}

// Get Ssp Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetSsp() string {
	return r._ssp
}

// Set is FeeType Setter
// 结算类型
func (r *YunosAppstoreOpenGetadsAPIRequest) SetFeeType(_feeType string) error {
	r._feeType = _feeType
	r.Set("fee_type", _feeType)
	return nil
}

// Get FeeType Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetFeeType() string {
	return r._feeType
}

// Set is ClientIp Setter
// 客户端来源ip
func (r *YunosAppstoreOpenGetadsAPIRequest) SetClientIp(_clientIp string) error {
	r._clientIp = _clientIp
	r.Set("client_ip", _clientIp)
	return nil
}

// Get ClientIp Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetClientIp() string {
	return r._clientIp
}

// Set is Pkgs Setter
// 广告指定包名
func (r *YunosAppstoreOpenGetadsAPIRequest) SetPkgs(_pkgs []string) error {
	r._pkgs = _pkgs
	r.Set("pkgs", _pkgs)
	return nil
}

// Get Pkgs Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetPkgs() []string {
	return r._pkgs
}

// Set is ClientVerCode Setter
// 客户端版本号
func (r *YunosAppstoreOpenGetadsAPIRequest) SetClientVerCode(_clientVerCode int64) error {
	r._clientVerCode = _clientVerCode
	r.Set("client_ver_code", _clientVerCode)
	return nil
}

// Get ClientVerCode Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetClientVerCode() int64 {
	return r._clientVerCode
}

// Set is TryMapToUuid Setter
// 是否映射到uuid
func (r *YunosAppstoreOpenGetadsAPIRequest) SetTryMapToUuid(_tryMapToUuid bool) error {
	r._tryMapToUuid = _tryMapToUuid
	r.Set("try_map_to_uuid", _tryMapToUuid)
	return nil
}

// Get TryMapToUuid Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetTryMapToUuid() bool {
	return r._tryMapToUuid
}

// Set is ExcludePkgs Setter
// 排除包名列表
func (r *YunosAppstoreOpenGetadsAPIRequest) SetExcludePkgs(_excludePkgs []string) error {
	r._excludePkgs = _excludePkgs
	r.Set("exclude_pkgs", _excludePkgs)
	return nil
}

// Get ExcludePkgs Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetExcludePkgs() []string {
	return r._excludePkgs
}

// Set is DeviceId Setter
// 设备唯一标识
func (r *YunosAppstoreOpenGetadsAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// Get DeviceId Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// Set is Size Setter
// 广告数量
func (r *YunosAppstoreOpenGetadsAPIRequest) SetSize(_size int64) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// Get Size Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetSize() int64 {
	return r._size
}

// Set is ExcludeCats Setter
// 排除分类
func (r *YunosAppstoreOpenGetadsAPIRequest) SetExcludeCats(_excludeCats []string) error {
	r._excludeCats = _excludeCats
	r.Set("exclude_cats", _excludeCats)
	return nil
}

// Get ExcludeCats Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetExcludeCats() []string {
	return r._excludeCats
}

// Set is TemplateIds Setter
// 创意模板id列表
func (r *YunosAppstoreOpenGetadsAPIRequest) SetTemplateIds(_templateIds []int64) error {
	r._templateIds = _templateIds
	r.Set("template_ids", _templateIds)
	return nil
}

// Get TemplateIds Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetTemplateIds() []int64 {
	return r._templateIds
}

// Set is Mrp Setter
// 广告底价
func (r *YunosAppstoreOpenGetadsAPIRequest) SetMrp(_mrp int64) error {
	r._mrp = _mrp
	r.Set("mrp", _mrp)
	return nil
}

// Get Mrp Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetMrp() int64 {
	return r._mrp
}

// Set is Options Setter
// 请求特征集
func (r *YunosAppstoreOpenGetadsAPIRequest) SetOptions(_options int64) error {
	r._options = _options
	r.Set("options", _options)
	return nil
}

// Get Options Getter
func (r YunosAppstoreOpenGetadsAPIRequest) GetOptions() int64 {
	return r._options
}
