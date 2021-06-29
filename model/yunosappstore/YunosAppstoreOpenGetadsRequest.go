package yunosappstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取外投广告 API请求
yunos.appstore.open.getads

将广告外投给外部合作伙伴
*/
type YunosAppstoreOpenGetadsRequest struct {
    model.Params
    // 请求id
    _rid   string
    // 指定广告分类
    _cats   []string
    // 是否排除已安装
    _excludeInstall   bool
    // 场景或页面标识
    _caseId   string
    // ssp标识
    _ssp   string
    // 结算类型
    _feeType   string
    // 客户端来源ip
    _clientIp   string
    // 广告指定包名
    _pkgs   []string
    // 客户端版本号
    _clientVerCode   int64
    // 是否映射到uuid
    _tryMapToUuid   bool
    // 排除包名列表
    _excludePkgs   []string
    // 设备唯一标识
    _deviceId   string
    // 广告数量
    _size   int64
    // 排除分类
    _excludeCats   []string
    // 创意模板id列表
    _templateIds   []int64
    // 广告底价
    _mrp   int64
    // 请求特征集
    _options   int64
}

// 初始化YunosAppstoreOpenGetadsRequest对象
func NewYunosAppstoreOpenGetadsRequest() *YunosAppstoreOpenGetadsRequest{
    return &YunosAppstoreOpenGetadsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosAppstoreOpenGetadsRequest) GetApiMethodName() string {
    return "yunos.appstore.open.getads"
}

// IRequest interface 方法, 获取API参数
func (r YunosAppstoreOpenGetadsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rid Setter
// 请求id
func (r *YunosAppstoreOpenGetadsRequest) SetRid(_rid string) error {
    r._rid = _rid
    r.Set("rid", _rid)
    return nil
}

// Rid Getter
func (r YunosAppstoreOpenGetadsRequest) GetRid() string {
    return r._rid
}
// Cats Setter
// 指定广告分类
func (r *YunosAppstoreOpenGetadsRequest) SetCats(_cats []string) error {
    r._cats = _cats
    r.Set("cats", _cats)
    return nil
}

// Cats Getter
func (r YunosAppstoreOpenGetadsRequest) GetCats() []string {
    return r._cats
}
// ExcludeInstall Setter
// 是否排除已安装
func (r *YunosAppstoreOpenGetadsRequest) SetExcludeInstall(_excludeInstall bool) error {
    r._excludeInstall = _excludeInstall
    r.Set("exclude_install", _excludeInstall)
    return nil
}

// ExcludeInstall Getter
func (r YunosAppstoreOpenGetadsRequest) GetExcludeInstall() bool {
    return r._excludeInstall
}
// CaseId Setter
// 场景或页面标识
func (r *YunosAppstoreOpenGetadsRequest) SetCaseId(_caseId string) error {
    r._caseId = _caseId
    r.Set("case_id", _caseId)
    return nil
}

// CaseId Getter
func (r YunosAppstoreOpenGetadsRequest) GetCaseId() string {
    return r._caseId
}
// Ssp Setter
// ssp标识
func (r *YunosAppstoreOpenGetadsRequest) SetSsp(_ssp string) error {
    r._ssp = _ssp
    r.Set("ssp", _ssp)
    return nil
}

// Ssp Getter
func (r YunosAppstoreOpenGetadsRequest) GetSsp() string {
    return r._ssp
}
// FeeType Setter
// 结算类型
func (r *YunosAppstoreOpenGetadsRequest) SetFeeType(_feeType string) error {
    r._feeType = _feeType
    r.Set("fee_type", _feeType)
    return nil
}

// FeeType Getter
func (r YunosAppstoreOpenGetadsRequest) GetFeeType() string {
    return r._feeType
}
// ClientIp Setter
// 客户端来源ip
func (r *YunosAppstoreOpenGetadsRequest) SetClientIp(_clientIp string) error {
    r._clientIp = _clientIp
    r.Set("client_ip", _clientIp)
    return nil
}

// ClientIp Getter
func (r YunosAppstoreOpenGetadsRequest) GetClientIp() string {
    return r._clientIp
}
// Pkgs Setter
// 广告指定包名
func (r *YunosAppstoreOpenGetadsRequest) SetPkgs(_pkgs []string) error {
    r._pkgs = _pkgs
    r.Set("pkgs", _pkgs)
    return nil
}

// Pkgs Getter
func (r YunosAppstoreOpenGetadsRequest) GetPkgs() []string {
    return r._pkgs
}
// ClientVerCode Setter
// 客户端版本号
func (r *YunosAppstoreOpenGetadsRequest) SetClientVerCode(_clientVerCode int64) error {
    r._clientVerCode = _clientVerCode
    r.Set("client_ver_code", _clientVerCode)
    return nil
}

// ClientVerCode Getter
func (r YunosAppstoreOpenGetadsRequest) GetClientVerCode() int64 {
    return r._clientVerCode
}
// TryMapToUuid Setter
// 是否映射到uuid
func (r *YunosAppstoreOpenGetadsRequest) SetTryMapToUuid(_tryMapToUuid bool) error {
    r._tryMapToUuid = _tryMapToUuid
    r.Set("try_map_to_uuid", _tryMapToUuid)
    return nil
}

// TryMapToUuid Getter
func (r YunosAppstoreOpenGetadsRequest) GetTryMapToUuid() bool {
    return r._tryMapToUuid
}
// ExcludePkgs Setter
// 排除包名列表
func (r *YunosAppstoreOpenGetadsRequest) SetExcludePkgs(_excludePkgs []string) error {
    r._excludePkgs = _excludePkgs
    r.Set("exclude_pkgs", _excludePkgs)
    return nil
}

// ExcludePkgs Getter
func (r YunosAppstoreOpenGetadsRequest) GetExcludePkgs() []string {
    return r._excludePkgs
}
// DeviceId Setter
// 设备唯一标识
func (r *YunosAppstoreOpenGetadsRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r YunosAppstoreOpenGetadsRequest) GetDeviceId() string {
    return r._deviceId
}
// Size Setter
// 广告数量
func (r *YunosAppstoreOpenGetadsRequest) SetSize(_size int64) error {
    r._size = _size
    r.Set("size", _size)
    return nil
}

// Size Getter
func (r YunosAppstoreOpenGetadsRequest) GetSize() int64 {
    return r._size
}
// ExcludeCats Setter
// 排除分类
func (r *YunosAppstoreOpenGetadsRequest) SetExcludeCats(_excludeCats []string) error {
    r._excludeCats = _excludeCats
    r.Set("exclude_cats", _excludeCats)
    return nil
}

// ExcludeCats Getter
func (r YunosAppstoreOpenGetadsRequest) GetExcludeCats() []string {
    return r._excludeCats
}
// TemplateIds Setter
// 创意模板id列表
func (r *YunosAppstoreOpenGetadsRequest) SetTemplateIds(_templateIds []int64) error {
    r._templateIds = _templateIds
    r.Set("template_ids", _templateIds)
    return nil
}

// TemplateIds Getter
func (r YunosAppstoreOpenGetadsRequest) GetTemplateIds() []int64 {
    return r._templateIds
}
// Mrp Setter
// 广告底价
func (r *YunosAppstoreOpenGetadsRequest) SetMrp(_mrp int64) error {
    r._mrp = _mrp
    r.Set("mrp", _mrp)
    return nil
}

// Mrp Getter
func (r YunosAppstoreOpenGetadsRequest) GetMrp() int64 {
    return r._mrp
}
// Options Setter
// 请求特征集
func (r *YunosAppstoreOpenGetadsRequest) SetOptions(_options int64) error {
    r._options = _options
    r.Set("options", _options)
    return nil
}

// Options Getter
func (r YunosAppstoreOpenGetadsRequest) GetOptions() int64 {
    return r._options
}
