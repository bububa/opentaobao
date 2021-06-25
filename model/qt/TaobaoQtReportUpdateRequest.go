package qt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新质检报告 APIRequest
taobao.qt.report.update

更新质检报告
*/
type TaobaoQtReportUpdateRequest struct {
    model.Params

    // 收费项code
    servcieItemCode   string 

    // 质检服务商名称
    spName   string 

    // 送检者昵称
    nick   string 

    // 一个质检服务唯一标识质量检验单的编号
    qtCode   string 

    // 质检名称
    qtName   string 

    // 宝贝样品url
    itemUrl   string 

    // 样品宝贝描述
    itemDesc   string 

    // (1L, "聚划算"),<br/>(2L, "消保"),<br/>(3L, "分销"),<br/>(4L, "抽检"),<br/>(5L, "良无限线下数据"),<br/>(6L, "入驻/续签商城"),<br/>(7L, "买家质检维权"),<br/>(8L, "实地验证"),<br/>(9L, "淘宝买家订单商品鉴定"),<br/>(10L,"假一赔三");
    qtType   int64 

    // 质检标准<br/>status=3 时必须非空
    qtStandard   string 

    // 质检报告源文件url<br/>status状态为3时必须非空
    reportUrl   string 

    // 0:已提交申请<br/>1:已收到样品<br/>2:已出检测结果<br/>3.已出具报告
    status   int64 

    // 只有status=3时赋值, <br/>0:未通过1:通过 空代表未判定
    isPassed   bool 

    // 检测结果消息描述
    message   string 

    // 自定义属性字段;分号分隔
    extAttr   string 

    // 送检日期
    gmtSubmit   string 

    // 提交报告结果时间
    gmtReport   string 

    // 质检有效到期时间，一般为一年有效期<br/>status状态为3时必须非空
    gmtExpiry   string 

    // 外部ID，和QT_TYPE 一起表示某种平台的实体ID。
QT_TYPE=9的时候，num_iid为淘宝订单号
    numIid   int64 

}

func NewTaobaoQtReportUpdateRequest() *TaobaoQtReportUpdateRequest{
    return &TaobaoQtReportUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQtReportUpdateRequest) GetApiMethodName() string {
    return "taobao.qt.report.update"
}

func (r TaobaoQtReportUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQtReportUpdateRequest) SetServcieItemCode(servcieItemCode string) error {
    r.servcieItemCode = servcieItemCode
    r.Set("servcie_item_code", servcieItemCode)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetServcieItemCode() string {
    return r.servcieItemCode
}

func (r *TaobaoQtReportUpdateRequest) SetSpName(spName string) error {
    r.spName = spName
    r.Set("sp_name", spName)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetSpName() string {
    return r.spName
}

func (r *TaobaoQtReportUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoQtReportUpdateRequest) SetQtCode(qtCode string) error {
    r.qtCode = qtCode
    r.Set("qt_code", qtCode)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetQtCode() string {
    return r.qtCode
}

func (r *TaobaoQtReportUpdateRequest) SetQtName(qtName string) error {
    r.qtName = qtName
    r.Set("qt_name", qtName)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetQtName() string {
    return r.qtName
}

func (r *TaobaoQtReportUpdateRequest) SetItemUrl(itemUrl string) error {
    r.itemUrl = itemUrl
    r.Set("item_url", itemUrl)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetItemUrl() string {
    return r.itemUrl
}

func (r *TaobaoQtReportUpdateRequest) SetItemDesc(itemDesc string) error {
    r.itemDesc = itemDesc
    r.Set("item_desc", itemDesc)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetItemDesc() string {
    return r.itemDesc
}

func (r *TaobaoQtReportUpdateRequest) SetQtType(qtType int64) error {
    r.qtType = qtType
    r.Set("qt_type", qtType)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetQtType() int64 {
    return r.qtType
}

func (r *TaobaoQtReportUpdateRequest) SetQtStandard(qtStandard string) error {
    r.qtStandard = qtStandard
    r.Set("qt_standard", qtStandard)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetQtStandard() string {
    return r.qtStandard
}

func (r *TaobaoQtReportUpdateRequest) SetReportUrl(reportUrl string) error {
    r.reportUrl = reportUrl
    r.Set("report_url", reportUrl)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetReportUrl() string {
    return r.reportUrl
}

func (r *TaobaoQtReportUpdateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetStatus() int64 {
    return r.status
}

func (r *TaobaoQtReportUpdateRequest) SetIsPassed(isPassed bool) error {
    r.isPassed = isPassed
    r.Set("is_passed", isPassed)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetIsPassed() bool {
    return r.isPassed
}

func (r *TaobaoQtReportUpdateRequest) SetMessage(message string) error {
    r.message = message
    r.Set("message", message)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetMessage() string {
    return r.message
}

func (r *TaobaoQtReportUpdateRequest) SetExtAttr(extAttr string) error {
    r.extAttr = extAttr
    r.Set("ext_attr", extAttr)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetExtAttr() string {
    return r.extAttr
}

func (r *TaobaoQtReportUpdateRequest) SetGmtSubmit(gmtSubmit string) error {
    r.gmtSubmit = gmtSubmit
    r.Set("gmt_submit", gmtSubmit)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetGmtSubmit() string {
    return r.gmtSubmit
}

func (r *TaobaoQtReportUpdateRequest) SetGmtReport(gmtReport string) error {
    r.gmtReport = gmtReport
    r.Set("gmt_report", gmtReport)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetGmtReport() string {
    return r.gmtReport
}

func (r *TaobaoQtReportUpdateRequest) SetGmtExpiry(gmtExpiry string) error {
    r.gmtExpiry = gmtExpiry
    r.Set("gmt_expiry", gmtExpiry)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetGmtExpiry() string {
    return r.gmtExpiry
}

func (r *TaobaoQtReportUpdateRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoQtReportUpdateRequest) GetNumIid() int64 {
    return r.numIid
}

