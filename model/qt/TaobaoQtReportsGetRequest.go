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
    spName   string
    // 质检类型，目前只支持查询qt_type=11的类型
    qtType   int64
    // 收费项code
    servcieItemCode   string
    // 送检者昵称
    nick   string
    // 查询时间段的开始时间
    startTime   string
    // 查询时间段的结束时间
    endTime   string
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
func (r *TaobaoQtReportsGetRequest) SetSpName(spName string) error {
    r.spName = spName
    r.Set("sp_name", spName)
    return nil
}

// SpName Getter
func (r TaobaoQtReportsGetRequest) GetSpName() string {
    return r.spName
}
// QtType Setter
// 质检类型，目前只支持查询qt_type=11的类型
func (r *TaobaoQtReportsGetRequest) SetQtType(qtType int64) error {
    r.qtType = qtType
    r.Set("qt_type", qtType)
    return nil
}

// QtType Getter
func (r TaobaoQtReportsGetRequest) GetQtType() int64 {
    return r.qtType
}
// ServcieItemCode Setter
// 收费项code
func (r *TaobaoQtReportsGetRequest) SetServcieItemCode(servcieItemCode string) error {
    r.servcieItemCode = servcieItemCode
    r.Set("servcie_item_code", servcieItemCode)
    return nil
}

// ServcieItemCode Getter
func (r TaobaoQtReportsGetRequest) GetServcieItemCode() string {
    return r.servcieItemCode
}
// Nick Setter
// 送检者昵称
func (r *TaobaoQtReportsGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoQtReportsGetRequest) GetNick() string {
    return r.nick
}
// StartTime Setter
// 查询时间段的开始时间
func (r *TaobaoQtReportsGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoQtReportsGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 查询时间段的结束时间
func (r *TaobaoQtReportsGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoQtReportsGetRequest) GetEndTime() string {
    return r.endTime
}
