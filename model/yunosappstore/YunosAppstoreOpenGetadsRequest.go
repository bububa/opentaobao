package yunosappstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取外投广告 APIRequest
yunos.appstore.open.getads

将广告外投给外部合作伙伴
*/
type YunosAppstoreOpenGetadsRequest struct {
    model.Params

    // 请求id
    rid   string 

    // 指定广告分类
    cats   []string 

    // 是否排除已安装
    excludeInstall   bool 

    // 场景或页面标识
    caseId   string 

    // ssp标识
    ssp   string 

    // 结算类型
    feeType   string 

    // 客户端来源ip
    clientIp   string 

    // 广告指定包名
    pkgs   []string 

    // 客户端版本号
    clientVerCode   int64 

    // 是否映射到uuid
    tryMapToUuid   bool 

    // 排除包名列表
    excludePkgs   []string 

    // 设备唯一标识
    deviceId   string 

    // 广告数量
    size   int64 

    // 排除分类
    excludeCats   []string 

    // 创意模板id列表
    templateIds   []int64 

    // 广告底价
    mrp   int64 

    // 请求特征集
    options   int64 

}

func NewYunosAppstoreOpenGetadsRequest() *YunosAppstoreOpenGetadsRequest{
    return &YunosAppstoreOpenGetadsRequest{
        Params: model.NewParams(),
    }
}

func (r YunosAppstoreOpenGetadsRequest) GetApiMethodName() string {
    return "yunos.appstore.open.getads"
}

func (r YunosAppstoreOpenGetadsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosAppstoreOpenGetadsRequest) SetRid(rid string) error {
    r.rid = rid
    r.Set("rid", rid)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetRid() string {
    return r.rid
}

func (r *YunosAppstoreOpenGetadsRequest) SetCats(cats []string) error {
    r.cats = cats
    r.Set("cats", cats)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetCats() []string {
    return r.cats
}

func (r *YunosAppstoreOpenGetadsRequest) SetExcludeInstall(excludeInstall bool) error {
    r.excludeInstall = excludeInstall
    r.Set("exclude_install", excludeInstall)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetExcludeInstall() bool {
    return r.excludeInstall
}

func (r *YunosAppstoreOpenGetadsRequest) SetCaseId(caseId string) error {
    r.caseId = caseId
    r.Set("case_id", caseId)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetCaseId() string {
    return r.caseId
}

func (r *YunosAppstoreOpenGetadsRequest) SetSsp(ssp string) error {
    r.ssp = ssp
    r.Set("ssp", ssp)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetSsp() string {
    return r.ssp
}

func (r *YunosAppstoreOpenGetadsRequest) SetFeeType(feeType string) error {
    r.feeType = feeType
    r.Set("fee_type", feeType)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetFeeType() string {
    return r.feeType
}

func (r *YunosAppstoreOpenGetadsRequest) SetClientIp(clientIp string) error {
    r.clientIp = clientIp
    r.Set("client_ip", clientIp)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetClientIp() string {
    return r.clientIp
}

func (r *YunosAppstoreOpenGetadsRequest) SetPkgs(pkgs []string) error {
    r.pkgs = pkgs
    r.Set("pkgs", pkgs)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetPkgs() []string {
    return r.pkgs
}

func (r *YunosAppstoreOpenGetadsRequest) SetClientVerCode(clientVerCode int64) error {
    r.clientVerCode = clientVerCode
    r.Set("client_ver_code", clientVerCode)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetClientVerCode() int64 {
    return r.clientVerCode
}

func (r *YunosAppstoreOpenGetadsRequest) SetTryMapToUuid(tryMapToUuid bool) error {
    r.tryMapToUuid = tryMapToUuid
    r.Set("try_map_to_uuid", tryMapToUuid)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetTryMapToUuid() bool {
    return r.tryMapToUuid
}

func (r *YunosAppstoreOpenGetadsRequest) SetExcludePkgs(excludePkgs []string) error {
    r.excludePkgs = excludePkgs
    r.Set("exclude_pkgs", excludePkgs)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetExcludePkgs() []string {
    return r.excludePkgs
}

func (r *YunosAppstoreOpenGetadsRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *YunosAppstoreOpenGetadsRequest) SetSize(size int64) error {
    r.size = size
    r.Set("size", size)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetSize() int64 {
    return r.size
}

func (r *YunosAppstoreOpenGetadsRequest) SetExcludeCats(excludeCats []string) error {
    r.excludeCats = excludeCats
    r.Set("exclude_cats", excludeCats)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetExcludeCats() []string {
    return r.excludeCats
}

func (r *YunosAppstoreOpenGetadsRequest) SetTemplateIds(templateIds []int64) error {
    r.templateIds = templateIds
    r.Set("template_ids", templateIds)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetTemplateIds() []int64 {
    return r.templateIds
}

func (r *YunosAppstoreOpenGetadsRequest) SetMrp(mrp int64) error {
    r.mrp = mrp
    r.Set("mrp", mrp)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetMrp() int64 {
    return r.mrp
}

func (r *YunosAppstoreOpenGetadsRequest) SetOptions(options int64) error {
    r.options = options
    r.Set("options", options)
    return nil
}

func (r YunosAppstoreOpenGetadsRequest) GetOptions() int64 {
    return r.options
}

