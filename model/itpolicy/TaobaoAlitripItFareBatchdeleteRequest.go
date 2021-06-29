package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票自有政策】批量删除 API请求
taobao.alitrip.it.fare.batchdelete

批量删除自有政策，单次删除最大5万，大于5万时候提示失败，需要缩小删除条件。此接口同步返回任务id，异步执行操作。每个接入方最多同时只能有10个处理中的任务，超过后直接返回失败。
*/
type TaobaoAlitripItFareBatchdeleteRequest struct {
    model.Params
    // 航空公司
    airline   string
    // 到达城市 可传多个 AND关系
    arrCity   string
    // 舱位 可传多个 或者的关系
    cabin   string
    // 是否能够混舱
    canRt   bool
    // 到达城市 可传多个 AND关系
    depCity   string
    // 最晚修改时间
    endModifyDate   string
    // 文件编号
    fileCode   string
    // 维护方式，可选值（UI：后台界面录入；EXCEL：后台excel批量导入；API：top接口添加）
    operateSource   string
    // 外部政策id
    outId   string
    // 最早修改时间
    startModifyDate   string
    // 去程适用开始日期
    startRestrictGoDate   string
    // 去程适用结束日期
    endRestrictGoDate   string
    // 运价类型，1单程 2往返
    fareType   int64
    // 0：未发布 1：已发布 2：已过期。不传的话，默认只能删除未发布和已过期的数据
    statusList   []int64
    // json格式的字符串，扩展属性，预留
    extendAttributes   string
}

// 初始化TaobaoAlitripItFareBatchdeleteRequest对象
func NewTaobaoAlitripItFareBatchdeleteRequest() *TaobaoAlitripItFareBatchdeleteRequest{
    return &TaobaoAlitripItFareBatchdeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItFareBatchdeleteRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.fare.batchdelete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItFareBatchdeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Airline Setter
// 航空公司
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetAirline(airline string) error {
    r.airline = airline
    r.Set("airline", airline)
    return nil
}

// Airline Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetAirline() string {
    return r.airline
}
// ArrCity Setter
// 到达城市 可传多个 AND关系
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetArrCity(arrCity string) error {
    r.arrCity = arrCity
    r.Set("arrCity", arrCity)
    return nil
}

// ArrCity Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetArrCity() string {
    return r.arrCity
}
// Cabin Setter
// 舱位 可传多个 或者的关系
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetCabin(cabin string) error {
    r.cabin = cabin
    r.Set("cabin", cabin)
    return nil
}

// Cabin Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetCabin() string {
    return r.cabin
}
// CanRt Setter
// 是否能够混舱
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetCanRt(canRt bool) error {
    r.canRt = canRt
    r.Set("canRt", canRt)
    return nil
}

// CanRt Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetCanRt() bool {
    return r.canRt
}
// DepCity Setter
// 到达城市 可传多个 AND关系
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetDepCity(depCity string) error {
    r.depCity = depCity
    r.Set("depCity", depCity)
    return nil
}

// DepCity Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetDepCity() string {
    return r.depCity
}
// EndModifyDate Setter
// 最晚修改时间
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetEndModifyDate(endModifyDate string) error {
    r.endModifyDate = endModifyDate
    r.Set("endModifyDate", endModifyDate)
    return nil
}

// EndModifyDate Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetEndModifyDate() string {
    return r.endModifyDate
}
// FileCode Setter
// 文件编号
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetFileCode(fileCode string) error {
    r.fileCode = fileCode
    r.Set("fileCode", fileCode)
    return nil
}

// FileCode Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetFileCode() string {
    return r.fileCode
}
// OperateSource Setter
// 维护方式，可选值（UI：后台界面录入；EXCEL：后台excel批量导入；API：top接口添加）
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetOperateSource(operateSource string) error {
    r.operateSource = operateSource
    r.Set("operateSource", operateSource)
    return nil
}

// OperateSource Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetOperateSource() string {
    return r.operateSource
}
// OutId Setter
// 外部政策id
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("outId", outId)
    return nil
}

// OutId Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetOutId() string {
    return r.outId
}
// StartModifyDate Setter
// 最早修改时间
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetStartModifyDate(startModifyDate string) error {
    r.startModifyDate = startModifyDate
    r.Set("startModifyDate", startModifyDate)
    return nil
}

// StartModifyDate Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetStartModifyDate() string {
    return r.startModifyDate
}
// StartRestrictGoDate Setter
// 去程适用开始日期
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetStartRestrictGoDate(startRestrictGoDate string) error {
    r.startRestrictGoDate = startRestrictGoDate
    r.Set("startRestrictGoDate", startRestrictGoDate)
    return nil
}

// StartRestrictGoDate Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetStartRestrictGoDate() string {
    return r.startRestrictGoDate
}
// EndRestrictGoDate Setter
// 去程适用结束日期
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetEndRestrictGoDate(endRestrictGoDate string) error {
    r.endRestrictGoDate = endRestrictGoDate
    r.Set("endRestrictGoDate", endRestrictGoDate)
    return nil
}

// EndRestrictGoDate Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetEndRestrictGoDate() string {
    return r.endRestrictGoDate
}
// FareType Setter
// 运价类型，1单程 2往返
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetFareType(fareType int64) error {
    r.fareType = fareType
    r.Set("fareType", fareType)
    return nil
}

// FareType Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetFareType() int64 {
    return r.fareType
}
// StatusList Setter
// 0：未发布 1：已发布 2：已过期。不传的话，默认只能删除未发布和已过期的数据
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetStatusList(statusList []int64) error {
    r.statusList = statusList
    r.Set("statusList", statusList)
    return nil
}

// StatusList Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetStatusList() []int64 {
    return r.statusList
}
// ExtendAttributes Setter
// json格式的字符串，扩展属性，预留
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetExtendAttributes(extendAttributes string) error {
    r.extendAttributes = extendAttributes
    r.Set("extendAttributes", extendAttributes)
    return nil
}

// ExtendAttributes Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetExtendAttributes() string {
    return r.extendAttributes
}
