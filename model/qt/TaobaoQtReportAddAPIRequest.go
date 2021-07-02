package qt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQtReportAddAPIRequest 上传质检报告 API请求
// taobao.qt.report.add
//
// 上传质检报告
type TaobaoQtReportAddAPIRequest struct {
	model.Params
	// 收费项code
	_servcieItemCode string
	// 质检服务商名称
	_spName string
	// 送检者昵称
	_nick string
	// 一个质检服务唯一标识质量检验单的编号
	_qtCode string
	// 质检名称
	_qtName string
	// 样品链接.<br/>QT_TYPE=9的时候，请填写N\A
	_itemUrl string
	// 样品信息描述
	_itemDesc string
	// (1L, "聚划算"),<br/>(2L, "消保"),<br/>(3L, "分销"),<br/>(4L, "抽检"),<br/>(5L, "良无限线下数据"),<br/>(6L, "入驻/续签商城"),<br/>(7L, "买家质检维权"),<br/>(8L, "实地验证"),<br/>(9L, "淘宝买家订单商品鉴定"),<br/>(10L,"假一赔三");
	_qtType int64
	// 质检标准<br/>status=3 时 必须非空
	_qtStandard string
	// 质检报告源文件url<br/>status状态为3时必须非空
	_reportUrl string
	// 0:已提交申请<br/>1:已收到样品<br/>2:已出检测结果<br/>3.已出具报告
	_status int64
	// 只有status=3时赋值, <br/>true 质检结果合格,false质检结果不合格.<br/>留空表示成分鉴定,不做判定
	_isPassed bool
	// 检测结果消息描述
	_message string
	// 自定义属性字段;分号分隔
	_extAttr string
	// 送检日期
	_gmtSubmit string
	// 提交报告结果时间
	_gmtReport string
	// 质检有效到期时间，一般为一年有效期<br/>status状态为3时必须非空
	_gmtExpiry string
	// 当前接口只有淘宝订单真假鉴定（QT_TYPE=9）的报告在该字段传入订单号，其他类型报告都不需要传输该值
	_numIid int64
}

// NewTaobaoQtReportAddRequest 初始化TaobaoQtReportAddAPIRequest对象
func NewTaobaoQtReportAddRequest() *TaobaoQtReportAddAPIRequest {
	return &TaobaoQtReportAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQtReportAddAPIRequest) GetApiMethodName() string {
	return "taobao.qt.report.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQtReportAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ServcieItemCode Setter
// 收费项code
func (r *TaobaoQtReportAddAPIRequest) SetServcieItemCode(_servcieItemCode string) error {
	r._servcieItemCode = _servcieItemCode
	r.Set("servcie_item_code", _servcieItemCode)
	return nil
}

// Get ServcieItemCode Getter
func (r TaobaoQtReportAddAPIRequest) GetServcieItemCode() string {
	return r._servcieItemCode
}

// Set is SpName Setter
// 质检服务商名称
func (r *TaobaoQtReportAddAPIRequest) SetSpName(_spName string) error {
	r._spName = _spName
	r.Set("sp_name", _spName)
	return nil
}

// Get SpName Getter
func (r TaobaoQtReportAddAPIRequest) GetSpName() string {
	return r._spName
}

// Set is Nick Setter
// 送检者昵称
func (r *TaobaoQtReportAddAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoQtReportAddAPIRequest) GetNick() string {
	return r._nick
}

// Set is QtCode Setter
// 一个质检服务唯一标识质量检验单的编号
func (r *TaobaoQtReportAddAPIRequest) SetQtCode(_qtCode string) error {
	r._qtCode = _qtCode
	r.Set("qt_code", _qtCode)
	return nil
}

// Get QtCode Getter
func (r TaobaoQtReportAddAPIRequest) GetQtCode() string {
	return r._qtCode
}

// Set is QtName Setter
// 质检名称
func (r *TaobaoQtReportAddAPIRequest) SetQtName(_qtName string) error {
	r._qtName = _qtName
	r.Set("qt_name", _qtName)
	return nil
}

// Get QtName Getter
func (r TaobaoQtReportAddAPIRequest) GetQtName() string {
	return r._qtName
}

// Set is ItemUrl Setter
// 样品链接.<br/>QT_TYPE=9的时候，请填写N\A
func (r *TaobaoQtReportAddAPIRequest) SetItemUrl(_itemUrl string) error {
	r._itemUrl = _itemUrl
	r.Set("item_url", _itemUrl)
	return nil
}

// Get ItemUrl Getter
func (r TaobaoQtReportAddAPIRequest) GetItemUrl() string {
	return r._itemUrl
}

// Set is ItemDesc Setter
// 样品信息描述
func (r *TaobaoQtReportAddAPIRequest) SetItemDesc(_itemDesc string) error {
	r._itemDesc = _itemDesc
	r.Set("item_desc", _itemDesc)
	return nil
}

// Get ItemDesc Getter
func (r TaobaoQtReportAddAPIRequest) GetItemDesc() string {
	return r._itemDesc
}

// Set is QtType Setter
// (1L, "聚划算"),<br/>(2L, "消保"),<br/>(3L, "分销"),<br/>(4L, "抽检"),<br/>(5L, "良无限线下数据"),<br/>(6L, "入驻/续签商城"),<br/>(7L, "买家质检维权"),<br/>(8L, "实地验证"),<br/>(9L, "淘宝买家订单商品鉴定"),<br/>(10L,"假一赔三");
func (r *TaobaoQtReportAddAPIRequest) SetQtType(_qtType int64) error {
	r._qtType = _qtType
	r.Set("qt_type", _qtType)
	return nil
}

// Get QtType Getter
func (r TaobaoQtReportAddAPIRequest) GetQtType() int64 {
	return r._qtType
}

// Set is QtStandard Setter
// 质检标准<br/>status=3 时 必须非空
func (r *TaobaoQtReportAddAPIRequest) SetQtStandard(_qtStandard string) error {
	r._qtStandard = _qtStandard
	r.Set("qt_standard", _qtStandard)
	return nil
}

// Get QtStandard Getter
func (r TaobaoQtReportAddAPIRequest) GetQtStandard() string {
	return r._qtStandard
}

// Set is ReportUrl Setter
// 质检报告源文件url<br/>status状态为3时必须非空
func (r *TaobaoQtReportAddAPIRequest) SetReportUrl(_reportUrl string) error {
	r._reportUrl = _reportUrl
	r.Set("report_url", _reportUrl)
	return nil
}

// Get ReportUrl Getter
func (r TaobaoQtReportAddAPIRequest) GetReportUrl() string {
	return r._reportUrl
}

// Set is Status Setter
// 0:已提交申请<br/>1:已收到样品<br/>2:已出检测结果<br/>3.已出具报告
func (r *TaobaoQtReportAddAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r TaobaoQtReportAddAPIRequest) GetStatus() int64 {
	return r._status
}

// Set is IsPassed Setter
// 只有status=3时赋值, <br/>true 质检结果合格,false质检结果不合格.<br/>留空表示成分鉴定,不做判定
func (r *TaobaoQtReportAddAPIRequest) SetIsPassed(_isPassed bool) error {
	r._isPassed = _isPassed
	r.Set("is_passed", _isPassed)
	return nil
}

// Get IsPassed Getter
func (r TaobaoQtReportAddAPIRequest) GetIsPassed() bool {
	return r._isPassed
}

// Set is Message Setter
// 检测结果消息描述
func (r *TaobaoQtReportAddAPIRequest) SetMessage(_message string) error {
	r._message = _message
	r.Set("message", _message)
	return nil
}

// Get Message Getter
func (r TaobaoQtReportAddAPIRequest) GetMessage() string {
	return r._message
}

// Set is ExtAttr Setter
// 自定义属性字段;分号分隔
func (r *TaobaoQtReportAddAPIRequest) SetExtAttr(_extAttr string) error {
	r._extAttr = _extAttr
	r.Set("ext_attr", _extAttr)
	return nil
}

// Get ExtAttr Getter
func (r TaobaoQtReportAddAPIRequest) GetExtAttr() string {
	return r._extAttr
}

// Set is GmtSubmit Setter
// 送检日期
func (r *TaobaoQtReportAddAPIRequest) SetGmtSubmit(_gmtSubmit string) error {
	r._gmtSubmit = _gmtSubmit
	r.Set("gmt_submit", _gmtSubmit)
	return nil
}

// Get GmtSubmit Getter
func (r TaobaoQtReportAddAPIRequest) GetGmtSubmit() string {
	return r._gmtSubmit
}

// Set is GmtReport Setter
// 提交报告结果时间
func (r *TaobaoQtReportAddAPIRequest) SetGmtReport(_gmtReport string) error {
	r._gmtReport = _gmtReport
	r.Set("gmt_report", _gmtReport)
	return nil
}

// Get GmtReport Getter
func (r TaobaoQtReportAddAPIRequest) GetGmtReport() string {
	return r._gmtReport
}

// Set is GmtExpiry Setter
// 质检有效到期时间，一般为一年有效期<br/>status状态为3时必须非空
func (r *TaobaoQtReportAddAPIRequest) SetGmtExpiry(_gmtExpiry string) error {
	r._gmtExpiry = _gmtExpiry
	r.Set("gmt_expiry", _gmtExpiry)
	return nil
}

// Get GmtExpiry Getter
func (r TaobaoQtReportAddAPIRequest) GetGmtExpiry() string {
	return r._gmtExpiry
}

// Set is NumIid Setter
// 当前接口只有淘宝订单真假鉴定（QT_TYPE=9）的报告在该字段传入订单号，其他类型报告都不需要传输该值
func (r *TaobaoQtReportAddAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// Get NumIid Getter
func (r TaobaoQtReportAddAPIRequest) GetNumIid() int64 {
	return r._numIid
}
