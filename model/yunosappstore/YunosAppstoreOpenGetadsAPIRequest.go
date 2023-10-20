package yunosappstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosappstoreopengetadsAPIRequest 获取外投广告 API请求
// yunos.appstore.open.getads
//
// 将广告外投给外部合作伙伴
type YunosappstoreopengetadsAPIRequest struct {
	model.Params
	// 指定广告分类
	_cats []string
	// 广告指定包名
	_pkgs []string
	// 排除包名列表
	_excludePkgs []string
	// 排除分类
	_excludeCats []string
	// 创意模板id列表
	_templateIds []int64
	// 场景或页面标识
	_caseId string
	// ssp标识
	_ssp string
	// 结算类型
	_feeType string
	// 客户端来源ip
	_clientIp string
	// 设备唯一标识
	_deviceId string
	// 请求id
	_rid string
	// 客户端版本号
	_clientVerCode int64
	// 广告数量
	_size int64
	// 广告底价
	_mrp int64
	// 请求特征集
	_options int64
	// 是否排除已安装
	_excludeInstall bool
	// 是否映射到uuid
	_tryMapToUuid bool
}

// NewYunosappstoreopengetadsRequest 初始化YunosappstoreopengetadsAPIRequest对象
func NewYunosappstoreopengetadsRequest() *YunosappstoreopengetadsAPIRequest {
	return &YunosappstoreopengetadsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosappstoreopengetadsAPIRequest) GetApiMethodName() string {
	return "yunos.appstore.open.getads"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosappstoreopengetadsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosappstoreopengetadsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCats is Cats Setter
// 指定广告分类
func (r *YunosappstoreopengetadsAPIRequest) SetCats(_cats []string) error {
	r._cats = _cats
	r.Set("cats", _cats)
	return nil
}

// GetCats Cats Getter
func (r YunosappstoreopengetadsAPIRequest) GetCats() []string {
	return r._cats
}

// SetPkgs is Pkgs Setter
// 广告指定包名
func (r *YunosappstoreopengetadsAPIRequest) SetPkgs(_pkgs []string) error {
	r._pkgs = _pkgs
	r.Set("pkgs", _pkgs)
	return nil
}

// GetPkgs Pkgs Getter
func (r YunosappstoreopengetadsAPIRequest) GetPkgs() []string {
	return r._pkgs
}

// SetExcludePkgs is ExcludePkgs Setter
// 排除包名列表
func (r *YunosappstoreopengetadsAPIRequest) SetExcludePkgs(_excludePkgs []string) error {
	r._excludePkgs = _excludePkgs
	r.Set("exclude_pkgs", _excludePkgs)
	return nil
}

// GetExcludePkgs ExcludePkgs Getter
func (r YunosappstoreopengetadsAPIRequest) GetExcludePkgs() []string {
	return r._excludePkgs
}

// SetExcludeCats is ExcludeCats Setter
// 排除分类
func (r *YunosappstoreopengetadsAPIRequest) SetExcludeCats(_excludeCats []string) error {
	r._excludeCats = _excludeCats
	r.Set("exclude_cats", _excludeCats)
	return nil
}

// GetExcludeCats ExcludeCats Getter
func (r YunosappstoreopengetadsAPIRequest) GetExcludeCats() []string {
	return r._excludeCats
}

// SetTemplateIds is TemplateIds Setter
// 创意模板id列表
func (r *YunosappstoreopengetadsAPIRequest) SetTemplateIds(_templateIds []int64) error {
	r._templateIds = _templateIds
	r.Set("template_ids", _templateIds)
	return nil
}

// GetTemplateIds TemplateIds Getter
func (r YunosappstoreopengetadsAPIRequest) GetTemplateIds() []int64 {
	return r._templateIds
}

// SetCaseId is CaseId Setter
// 场景或页面标识
func (r *YunosappstoreopengetadsAPIRequest) SetCaseId(_caseId string) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// GetCaseId CaseId Getter
func (r YunosappstoreopengetadsAPIRequest) GetCaseId() string {
	return r._caseId
}

// SetSsp is Ssp Setter
// ssp标识
func (r *YunosappstoreopengetadsAPIRequest) SetSsp(_ssp string) error {
	r._ssp = _ssp
	r.Set("ssp", _ssp)
	return nil
}

// GetSsp Ssp Getter
func (r YunosappstoreopengetadsAPIRequest) GetSsp() string {
	return r._ssp
}

// SetFeeType is FeeType Setter
// 结算类型
func (r *YunosappstoreopengetadsAPIRequest) SetFeeType(_feeType string) error {
	r._feeType = _feeType
	r.Set("fee_type", _feeType)
	return nil
}

// GetFeeType FeeType Getter
func (r YunosappstoreopengetadsAPIRequest) GetFeeType() string {
	return r._feeType
}

// SetClientIp is ClientIp Setter
// 客户端来源ip
func (r *YunosappstoreopengetadsAPIRequest) SetClientIp(_clientIp string) error {
	r._clientIp = _clientIp
	r.Set("client_ip", _clientIp)
	return nil
}

// GetClientIp ClientIp Getter
func (r YunosappstoreopengetadsAPIRequest) GetClientIp() string {
	return r._clientIp
}

// SetDeviceId is DeviceId Setter
// 设备唯一标识
func (r *YunosappstoreopengetadsAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r YunosappstoreopengetadsAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetRid is Rid Setter
// 请求id
func (r *YunosappstoreopengetadsAPIRequest) SetRid(_rid string) error {
	r._rid = _rid
	r.Set("rid", _rid)
	return nil
}

// GetRid Rid Getter
func (r YunosappstoreopengetadsAPIRequest) GetRid() string {
	return r._rid
}

// SetClientVerCode is ClientVerCode Setter
// 客户端版本号
func (r *YunosappstoreopengetadsAPIRequest) SetClientVerCode(_clientVerCode int64) error {
	r._clientVerCode = _clientVerCode
	r.Set("client_ver_code", _clientVerCode)
	return nil
}

// GetClientVerCode ClientVerCode Getter
func (r YunosappstoreopengetadsAPIRequest) GetClientVerCode() int64 {
	return r._clientVerCode
}

// SetSize is Size Setter
// 广告数量
func (r *YunosappstoreopengetadsAPIRequest) SetSize(_size int64) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// GetSize Size Getter
func (r YunosappstoreopengetadsAPIRequest) GetSize() int64 {
	return r._size
}

// SetMrp is Mrp Setter
// 广告底价
func (r *YunosappstoreopengetadsAPIRequest) SetMrp(_mrp int64) error {
	r._mrp = _mrp
	r.Set("mrp", _mrp)
	return nil
}

// GetMrp Mrp Getter
func (r YunosappstoreopengetadsAPIRequest) GetMrp() int64 {
	return r._mrp
}

// SetOptions is Options Setter
// 请求特征集
func (r *YunosappstoreopengetadsAPIRequest) SetOptions(_options int64) error {
	r._options = _options
	r.Set("options", _options)
	return nil
}

// GetOptions Options Getter
func (r YunosappstoreopengetadsAPIRequest) GetOptions() int64 {
	return r._options
}

// SetExcludeInstall is ExcludeInstall Setter
// 是否排除已安装
func (r *YunosappstoreopengetadsAPIRequest) SetExcludeInstall(_excludeInstall bool) error {
	r._excludeInstall = _excludeInstall
	r.Set("exclude_install", _excludeInstall)
	return nil
}

// GetExcludeInstall ExcludeInstall Getter
func (r YunosappstoreopengetadsAPIRequest) GetExcludeInstall() bool {
	return r._excludeInstall
}

// SetTryMapToUuid is TryMapToUuid Setter
// 是否映射到uuid
func (r *YunosappstoreopengetadsAPIRequest) SetTryMapToUuid(_tryMapToUuid bool) error {
	r._tryMapToUuid = _tryMapToUuid
	r.Set("try_map_to_uuid", _tryMapToUuid)
	return nil
}

// GetTryMapToUuid TryMapToUuid Getter
func (r YunosappstoreopengetadsAPIRequest) GetTryMapToUuid() bool {
	return r._tryMapToUuid
}
