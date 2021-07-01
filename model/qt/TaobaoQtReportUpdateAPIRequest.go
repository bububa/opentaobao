package qt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新质检报告 API请求
taobao.qt.report.update

更新质检报告
*/
type TaobaoQtReportUpdateAPIRequest struct {
    model.Params
    // 收费项code
    _servcieItemCode   string
    // 质检服务商名称
    _spName   string
    // 送检者昵称
    _nick   string
    // 一个质检服务唯一标识质量检验单的编号
    _qtCode   string
    // 质检名称
    _qtName   string
    // 宝贝样品url
    _itemUrl   string
    // 样品宝贝描述
    _itemDesc   string
    // (1L, "聚划算"),<br/>(2L, "消保"),<br/>(3L, "分销"),<br/>(4L, "抽检"),<br/>(5L, "良无限线下数据"),<br/>(6L, "入驻/续签商城"),<br/>(7L, "买家质检维权"),<br/>(8L, "实地验证"),<br/>(9L, "淘宝买家订单商品鉴定"),<br/>(10L,"假一赔三");
    _qtType   int64
    // 质检标准<br/>status=3 时必须非空
    _qtStandard   string
    // 质检报告源文件url<br/>status状态为3时必须非空
    _reportUrl   string
    // 0:已提交申请<br/>1:已收到样品<br/>2:已出检测结果<br/>3.已出具报告
    _status   int64
    // 只有status=3时赋值, <br/>0:未通过1:通过 空代表未判定
    _isPassed   bool
    // 检测结果消息描述
    _message   string
    // 自定义属性字段;分号分隔
    _extAttr   string
    // 送检日期
    _gmtSubmit   string
    // 提交报告结果时间
    _gmtReport   string
    // 质检有效到期时间，一般为一年有效期<br/>status状态为3时必须非空
    _gmtExpiry   string
    // 外部ID，和QT_TYPE 一起表示某种平台的实体ID。QT_TYPE=9的时候，num_iid为淘宝订单号
    _numIid   int64
}

// 初始化TaobaoQtReportUpdateAPIRequest对象
func NewTaobaoQtReportUpdateRequest() *TaobaoQtReportUpdateAPIRequest{
    return &TaobaoQtReportUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQtReportUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.qt.report.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQtReportUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServcieItemCode Setter
// 收费项code
func (r *TaobaoQtReportUpdateAPIRequest) SetServcieItemCode(_servcieItemCode string) error {
    r._servcieItemCode = _servcieItemCode
    r.Set("servcie_item_code", _servcieItemCode)
    return nil
}

// ServcieItemCode Getter
func (r TaobaoQtReportUpdateAPIRequest) GetServcieItemCode() string {
    return r._servcieItemCode
}
// SpName Setter
// 质检服务商名称
func (r *TaobaoQtReportUpdateAPIRequest) SetSpName(_spName string) error {
    r._spName = _spName
    r.Set("sp_name", _spName)
    return nil
}

// SpName Getter
func (r TaobaoQtReportUpdateAPIRequest) GetSpName() string {
    return r._spName
}
// Nick Setter
// 送检者昵称
func (r *TaobaoQtReportUpdateAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoQtReportUpdateAPIRequest) GetNick() string {
    return r._nick
}
// QtCode Setter
// 一个质检服务唯一标识质量检验单的编号
func (r *TaobaoQtReportUpdateAPIRequest) SetQtCode(_qtCode string) error {
    r._qtCode = _qtCode
    r.Set("qt_code", _qtCode)
    return nil
}

// QtCode Getter
func (r TaobaoQtReportUpdateAPIRequest) GetQtCode() string {
    return r._qtCode
}
// QtName Setter
// 质检名称
func (r *TaobaoQtReportUpdateAPIRequest) SetQtName(_qtName string) error {
    r._qtName = _qtName
    r.Set("qt_name", _qtName)
    return nil
}

// QtName Getter
func (r TaobaoQtReportUpdateAPIRequest) GetQtName() string {
    return r._qtName
}
// ItemUrl Setter
// 宝贝样品url
func (r *TaobaoQtReportUpdateAPIRequest) SetItemUrl(_itemUrl string) error {
    r._itemUrl = _itemUrl
    r.Set("item_url", _itemUrl)
    return nil
}

// ItemUrl Getter
func (r TaobaoQtReportUpdateAPIRequest) GetItemUrl() string {
    return r._itemUrl
}
// ItemDesc Setter
// 样品宝贝描述
func (r *TaobaoQtReportUpdateAPIRequest) SetItemDesc(_itemDesc string) error {
    r._itemDesc = _itemDesc
    r.Set("item_desc", _itemDesc)
    return nil
}

// ItemDesc Getter
func (r TaobaoQtReportUpdateAPIRequest) GetItemDesc() string {
    return r._itemDesc
}
// QtType Setter
// (1L, "聚划算"),<br/>(2L, "消保"),<br/>(3L, "分销"),<br/>(4L, "抽检"),<br/>(5L, "良无限线下数据"),<br/>(6L, "入驻/续签商城"),<br/>(7L, "买家质检维权"),<br/>(8L, "实地验证"),<br/>(9L, "淘宝买家订单商品鉴定"),<br/>(10L,"假一赔三");
func (r *TaobaoQtReportUpdateAPIRequest) SetQtType(_qtType int64) error {
    r._qtType = _qtType
    r.Set("qt_type", _qtType)
    return nil
}

// QtType Getter
func (r TaobaoQtReportUpdateAPIRequest) GetQtType() int64 {
    return r._qtType
}
// QtStandard Setter
// 质检标准<br/>status=3 时必须非空
func (r *TaobaoQtReportUpdateAPIRequest) SetQtStandard(_qtStandard string) error {
    r._qtStandard = _qtStandard
    r.Set("qt_standard", _qtStandard)
    return nil
}

// QtStandard Getter
func (r TaobaoQtReportUpdateAPIRequest) GetQtStandard() string {
    return r._qtStandard
}
// ReportUrl Setter
// 质检报告源文件url<br/>status状态为3时必须非空
func (r *TaobaoQtReportUpdateAPIRequest) SetReportUrl(_reportUrl string) error {
    r._reportUrl = _reportUrl
    r.Set("report_url", _reportUrl)
    return nil
}

// ReportUrl Getter
func (r TaobaoQtReportUpdateAPIRequest) GetReportUrl() string {
    return r._reportUrl
}
// Status Setter
// 0:已提交申请<br/>1:已收到样品<br/>2:已出检测结果<br/>3.已出具报告
func (r *TaobaoQtReportUpdateAPIRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoQtReportUpdateAPIRequest) GetStatus() int64 {
    return r._status
}
// IsPassed Setter
// 只有status=3时赋值, <br/>0:未通过1:通过 空代表未判定
func (r *TaobaoQtReportUpdateAPIRequest) SetIsPassed(_isPassed bool) error {
    r._isPassed = _isPassed
    r.Set("is_passed", _isPassed)
    return nil
}

// IsPassed Getter
func (r TaobaoQtReportUpdateAPIRequest) GetIsPassed() bool {
    return r._isPassed
}
// Message Setter
// 检测结果消息描述
func (r *TaobaoQtReportUpdateAPIRequest) SetMessage(_message string) error {
    r._message = _message
    r.Set("message", _message)
    return nil
}

// Message Getter
func (r TaobaoQtReportUpdateAPIRequest) GetMessage() string {
    return r._message
}
// ExtAttr Setter
// 自定义属性字段;分号分隔
func (r *TaobaoQtReportUpdateAPIRequest) SetExtAttr(_extAttr string) error {
    r._extAttr = _extAttr
    r.Set("ext_attr", _extAttr)
    return nil
}

// ExtAttr Getter
func (r TaobaoQtReportUpdateAPIRequest) GetExtAttr() string {
    return r._extAttr
}
// GmtSubmit Setter
// 送检日期
func (r *TaobaoQtReportUpdateAPIRequest) SetGmtSubmit(_gmtSubmit string) error {
    r._gmtSubmit = _gmtSubmit
    r.Set("gmt_submit", _gmtSubmit)
    return nil
}

// GmtSubmit Getter
func (r TaobaoQtReportUpdateAPIRequest) GetGmtSubmit() string {
    return r._gmtSubmit
}
// GmtReport Setter
// 提交报告结果时间
func (r *TaobaoQtReportUpdateAPIRequest) SetGmtReport(_gmtReport string) error {
    r._gmtReport = _gmtReport
    r.Set("gmt_report", _gmtReport)
    return nil
}

// GmtReport Getter
func (r TaobaoQtReportUpdateAPIRequest) GetGmtReport() string {
    return r._gmtReport
}
// GmtExpiry Setter
// 质检有效到期时间，一般为一年有效期<br/>status状态为3时必须非空
func (r *TaobaoQtReportUpdateAPIRequest) SetGmtExpiry(_gmtExpiry string) error {
    r._gmtExpiry = _gmtExpiry
    r.Set("gmt_expiry", _gmtExpiry)
    return nil
}

// GmtExpiry Getter
func (r TaobaoQtReportUpdateAPIRequest) GetGmtExpiry() string {
    return r._gmtExpiry
}
// NumIid Setter
// 外部ID，和QT_TYPE 一起表示某种平台的实体ID。QT_TYPE=9的时候，num_iid为淘宝订单号
func (r *TaobaoQtReportUpdateAPIRequest) SetNumIid(_numIid int64) error {
    r._numIid = _numIid
    r.Set("num_iid", _numIid)
    return nil
}

// NumIid Getter
func (r TaobaoQtReportUpdateAPIRequest) GetNumIid() int64 {
    return r._numIid
}
