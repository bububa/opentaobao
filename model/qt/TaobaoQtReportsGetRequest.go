package qt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询质检报告 API请求
taobao.qt.reports.get

批量查询质检报告，目前只支持查询qtType=11（天猫真假鉴定）类型的报告
*/
type TaobaoQtReportsGetRequest struct {
    model.Params
    // 质检服务商名
    _spName   string
    // 质检类型，目前只支持查询qt_type=11的类型
    _qtType   int64
    // 收费项code
    _servcieItemCode   string
    // 送检者昵称
    _nick   string
    // 查询时间段的开始时间
    _startTime   string
    // 查询时间段的结束时间
    _endTime   string
}

// 初始化TaobaoQtReportsGetRequest对象
func NewTaobaoQtReportsGetRequest() *TaobaoQtReportsGetRequest{
    return &TaobaoQtReportsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQtReportsGetRequest) GetApiMethodName() string {
    return "taobao.qt.reports.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQtReportsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SpName Setter
// 质检服务商名
func (r *TaobaoQtReportsGetRequest) SetSpName(_spName string) error {
    r._spName = _spName
    r.Set("sp_name", _spName)
    return nil
}

// SpName Getter
func (r TaobaoQtReportsGetRequest) GetSpName() string {
    return r._spName
}
// QtType Setter
// 质检类型，目前只支持查询qt_type=11的类型
func (r *TaobaoQtReportsGetRequest) SetQtType(_qtType int64) error {
    r._qtType = _qtType
    r.Set("qt_type", _qtType)
    return nil
}

// QtType Getter
func (r TaobaoQtReportsGetRequest) GetQtType() int64 {
    return r._qtType
}
// ServcieItemCode Setter
// 收费项code
func (r *TaobaoQtReportsGetRequest) SetServcieItemCode(_servcieItemCode string) error {
    r._servcieItemCode = _servcieItemCode
    r.Set("servcie_item_code", _servcieItemCode)
    return nil
}

// ServcieItemCode Getter
func (r TaobaoQtReportsGetRequest) GetServcieItemCode() string {
    return r._servcieItemCode
}
// Nick Setter
// 送检者昵称
func (r *TaobaoQtReportsGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoQtReportsGetRequest) GetNick() string {
    return r._nick
}
// StartTime Setter
// 查询时间段的开始时间
func (r *TaobaoQtReportsGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoQtReportsGetRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 查询时间段的结束时间
func (r *TaobaoQtReportsGetRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoQtReportsGetRequest) GetEndTime() string {
    return r._endTime
}
