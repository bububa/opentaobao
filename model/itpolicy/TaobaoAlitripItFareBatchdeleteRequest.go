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
    _airline   string
    // 到达城市 可传多个 AND关系
    _arrCity   string
    // 舱位 可传多个 或者的关系
    _cabin   string
    // 是否能够混舱
    _canRt   bool
    // 到达城市 可传多个 AND关系
    _depCity   string
    // 最晚修改时间
    _endModifyDate   string
    // 文件编号
    _fileCode   string
    // 维护方式，可选值（UI：后台界面录入；EXCEL：后台excel批量导入；API：top接口添加）
    _operateSource   string
    // 外部政策id
    _outId   string
    // 最早修改时间
    _startModifyDate   string
    // 去程适用开始日期
    _startRestrictGoDate   string
    // 去程适用结束日期
    _endRestrictGoDate   string
    // 运价类型，1单程 2往返
    _fareType   int64
    // 0：未发布 1：已发布 2：已过期。不传的话，默认只能删除未发布和已过期的数据
    _statusList   []int64
    // json格式的字符串，扩展属性，预留
    _extendAttributes   string
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
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetAirline(_airline string) error {
    r._airline = _airline
    r.Set("airline", _airline)
    return nil
}

// Airline Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetAirline() string {
    return r._airline
}
// ArrCity Setter
// 到达城市 可传多个 AND关系
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetArrCity(_arrCity string) error {
    r._arrCity = _arrCity
    r.Set("arrCity", _arrCity)
    return nil
}

// ArrCity Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetArrCity() string {
    return r._arrCity
}
// Cabin Setter
// 舱位 可传多个 或者的关系
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetCabin(_cabin string) error {
    r._cabin = _cabin
    r.Set("cabin", _cabin)
    return nil
}

// Cabin Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetCabin() string {
    return r._cabin
}
// CanRt Setter
// 是否能够混舱
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetCanRt(_canRt bool) error {
    r._canRt = _canRt
    r.Set("canRt", _canRt)
    return nil
}

// CanRt Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetCanRt() bool {
    return r._canRt
}
// DepCity Setter
// 到达城市 可传多个 AND关系
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetDepCity(_depCity string) error {
    r._depCity = _depCity
    r.Set("depCity", _depCity)
    return nil
}

// DepCity Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetDepCity() string {
    return r._depCity
}
// EndModifyDate Setter
// 最晚修改时间
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetEndModifyDate(_endModifyDate string) error {
    r._endModifyDate = _endModifyDate
    r.Set("endModifyDate", _endModifyDate)
    return nil
}

// EndModifyDate Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetEndModifyDate() string {
    return r._endModifyDate
}
// FileCode Setter
// 文件编号
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetFileCode(_fileCode string) error {
    r._fileCode = _fileCode
    r.Set("fileCode", _fileCode)
    return nil
}

// FileCode Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetFileCode() string {
    return r._fileCode
}
// OperateSource Setter
// 维护方式，可选值（UI：后台界面录入；EXCEL：后台excel批量导入；API：top接口添加）
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetOperateSource(_operateSource string) error {
    r._operateSource = _operateSource
    r.Set("operateSource", _operateSource)
    return nil
}

// OperateSource Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetOperateSource() string {
    return r._operateSource
}
// OutId Setter
// 外部政策id
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetOutId(_outId string) error {
    r._outId = _outId
    r.Set("outId", _outId)
    return nil
}

// OutId Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetOutId() string {
    return r._outId
}
// StartModifyDate Setter
// 最早修改时间
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetStartModifyDate(_startModifyDate string) error {
    r._startModifyDate = _startModifyDate
    r.Set("startModifyDate", _startModifyDate)
    return nil
}

// StartModifyDate Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetStartModifyDate() string {
    return r._startModifyDate
}
// StartRestrictGoDate Setter
// 去程适用开始日期
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetStartRestrictGoDate(_startRestrictGoDate string) error {
    r._startRestrictGoDate = _startRestrictGoDate
    r.Set("startRestrictGoDate", _startRestrictGoDate)
    return nil
}

// StartRestrictGoDate Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetStartRestrictGoDate() string {
    return r._startRestrictGoDate
}
// EndRestrictGoDate Setter
// 去程适用结束日期
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetEndRestrictGoDate(_endRestrictGoDate string) error {
    r._endRestrictGoDate = _endRestrictGoDate
    r.Set("endRestrictGoDate", _endRestrictGoDate)
    return nil
}

// EndRestrictGoDate Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetEndRestrictGoDate() string {
    return r._endRestrictGoDate
}
// FareType Setter
// 运价类型，1单程 2往返
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetFareType(_fareType int64) error {
    r._fareType = _fareType
    r.Set("fareType", _fareType)
    return nil
}

// FareType Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetFareType() int64 {
    return r._fareType
}
// StatusList Setter
// 0：未发布 1：已发布 2：已过期。不传的话，默认只能删除未发布和已过期的数据
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetStatusList(_statusList []int64) error {
    r._statusList = _statusList
    r.Set("statusList", _statusList)
    return nil
}

// StatusList Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetStatusList() []int64 {
    return r._statusList
}
// ExtendAttributes Setter
// json格式的字符串，扩展属性，预留
func (r *TaobaoAlitripItFareBatchdeleteRequest) SetExtendAttributes(_extendAttributes string) error {
    r._extendAttributes = _extendAttributes
    r.Set("extendAttributes", _extendAttributes)
    return nil
}

// ExtendAttributes Getter
func (r TaobaoAlitripItFareBatchdeleteRequest) GetExtendAttributes() string {
    return r._extendAttributes
}
