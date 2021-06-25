package qt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传质检报告 APIRequest
taobao.qt.report.add

上传质检报告
*/
type TaobaoQtReportAddRequest struct {
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

    // 样品链接.<br/>QT_TYPE=9的时候，请填写N\A
    itemUrl   string 

    // 样品信息描述
    itemDesc   string 

    // (1L, "聚划算"),<br/>(2L, "消保"),<br/>(3L, "分销"),<br/>(4L, "抽检"),<br/>(5L, "良无限线下数据"),<br/>(6L, "入驻/续签商城"),<br/>(7L, "买家质检维权"),<br/>(8L, "实地验证"),<br/>(9L, "淘宝买家订单商品鉴定"),<br/>(10L,"假一赔三");
    qtType   int64 

    // 质检标准<br/>status=3 时 必须非空
    qtStandard   string 

    // 质检报告源文件url<br/>status状态为3时必须非空
    reportUrl   string 

    // 0:已提交申请<br/>1:已收到样品<br/>2:已出检测结果<br/>3.已出具报告
    status   int64 

    // 只有status=3时赋值, <br/>true 质检结果合格,false质检结果不合格.<br/>留空表示成分鉴定,不做判定
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

    // 当前接口只有淘宝订单真假鉴定（QT_TYPE=9）的报告在该字段传入订单号，其他类型报告都不需要传输该值
    numIid   int64 

}

func NewTaobaoQtReportAddRequest() *TaobaoQtReportAddRequest{
    return &TaobaoQtReportAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQtReportAddRequest) GetApiMethodName() string {
    return "taobao.qt.report.add"
}

func (r TaobaoQtReportAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQtReportAddRequest) SetServcieItemCode(servcieItemCode string) error {
    r.servcieItemCode = servcieItemCode
    r.Set("servcie_item_code", servcieItemCode)
    return nil
}

func (r TaobaoQtReportAddRequest) GetServcieItemCode() string {
    return r.servcieItemCode
}

func (r *TaobaoQtReportAddRequest) SetSpName(spName string) error {
    r.spName = spName
    r.Set("sp_name", spName)
    return nil
}

func (r TaobaoQtReportAddRequest) GetSpName() string {
    return r.spName
}

func (r *TaobaoQtReportAddRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoQtReportAddRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoQtReportAddRequest) SetQtCode(qtCode string) error {
    r.qtCode = qtCode
    r.Set("qt_code", qtCode)
    return nil
}

func (r TaobaoQtReportAddRequest) GetQtCode() string {
    return r.qtCode
}

func (r *TaobaoQtReportAddRequest) SetQtName(qtName string) error {
    r.qtName = qtName
    r.Set("qt_name", qtName)
    return nil
}

func (r TaobaoQtReportAddRequest) GetQtName() string {
    return r.qtName
}

func (r *TaobaoQtReportAddRequest) SetItemUrl(itemUrl string) error {
    r.itemUrl = itemUrl
    r.Set("item_url", itemUrl)
    return nil
}

func (r TaobaoQtReportAddRequest) GetItemUrl() string {
    return r.itemUrl
}

func (r *TaobaoQtReportAddRequest) SetItemDesc(itemDesc string) error {
    r.itemDesc = itemDesc
    r.Set("item_desc", itemDesc)
    return nil
}

func (r TaobaoQtReportAddRequest) GetItemDesc() string {
    return r.itemDesc
}

func (r *TaobaoQtReportAddRequest) SetQtType(qtType int64) error {
    r.qtType = qtType
    r.Set("qt_type", qtType)
    return nil
}

func (r TaobaoQtReportAddRequest) GetQtType() int64 {
    return r.qtType
}

func (r *TaobaoQtReportAddRequest) SetQtStandard(qtStandard string) error {
    r.qtStandard = qtStandard
    r.Set("qt_standard", qtStandard)
    return nil
}

func (r TaobaoQtReportAddRequest) GetQtStandard() string {
    return r.qtStandard
}

func (r *TaobaoQtReportAddRequest) SetReportUrl(reportUrl string) error {
    r.reportUrl = reportUrl
    r.Set("report_url", reportUrl)
    return nil
}

func (r TaobaoQtReportAddRequest) GetReportUrl() string {
    return r.reportUrl
}

func (r *TaobaoQtReportAddRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoQtReportAddRequest) GetStatus() int64 {
    return r.status
}

func (r *TaobaoQtReportAddRequest) SetIsPassed(isPassed bool) error {
    r.isPassed = isPassed
    r.Set("is_passed", isPassed)
    return nil
}

func (r TaobaoQtReportAddRequest) GetIsPassed() bool {
    return r.isPassed
}

func (r *TaobaoQtReportAddRequest) SetMessage(message string) error {
    r.message = message
    r.Set("message", message)
    return nil
}

func (r TaobaoQtReportAddRequest) GetMessage() string {
    return r.message
}

func (r *TaobaoQtReportAddRequest) SetExtAttr(extAttr string) error {
    r.extAttr = extAttr
    r.Set("ext_attr", extAttr)
    return nil
}

func (r TaobaoQtReportAddRequest) GetExtAttr() string {
    return r.extAttr
}

func (r *TaobaoQtReportAddRequest) SetGmtSubmit(gmtSubmit string) error {
    r.gmtSubmit = gmtSubmit
    r.Set("gmt_submit", gmtSubmit)
    return nil
}

func (r TaobaoQtReportAddRequest) GetGmtSubmit() string {
    return r.gmtSubmit
}

func (r *TaobaoQtReportAddRequest) SetGmtReport(gmtReport string) error {
    r.gmtReport = gmtReport
    r.Set("gmt_report", gmtReport)
    return nil
}

func (r TaobaoQtReportAddRequest) GetGmtReport() string {
    return r.gmtReport
}

func (r *TaobaoQtReportAddRequest) SetGmtExpiry(gmtExpiry string) error {
    r.gmtExpiry = gmtExpiry
    r.Set("gmt_expiry", gmtExpiry)
    return nil
}

func (r TaobaoQtReportAddRequest) GetGmtExpiry() string {
    return r.gmtExpiry
}

func (r *TaobaoQtReportAddRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoQtReportAddRequest) GetNumIid() int64 {
    return r.numIid
}

