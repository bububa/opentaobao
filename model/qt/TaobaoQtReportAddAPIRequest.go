package qt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqtreportaddAPIRequest 上传质检报告 API请求
// taobao.qt.report.add
//
// 上传质检报告
type TaobaoqtreportaddAPIRequest struct {
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
	// 质检标准<br/>status=3 时 必须非空
	_qtStandard string
	// 质检报告源文件url<br/>status状态为3时必须非空
	_reportUrl string
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
	// (1L, "聚划算"),<br/>(2L, "消保"),<br/>(3L, "分销"),<br/>(4L, "抽检"),<br/>(5L, "良无限线下数据"),<br/>(6L, "入驻/续签商城"),<br/>(7L, "买家质检维权"),<br/>(8L, "实地验证"),<br/>(9L, "淘宝买家订单商品鉴定"),<br/>(10L,"假一赔三");
	_qtType int64
	// 0:已提交申请<br/>1:已收到样品<br/>2:已出检测结果<br/>3.已出具报告
	_status int64
	// 当前接口只有淘宝订单真假鉴定（QT_TYPE=9）的报告在该字段传入订单号，其他类型报告都不需要传输该值
	_numIid int64
	// 只有status=3时赋值, <br/>true 质检结果合格,false质检结果不合格.<br/>留空表示成分鉴定,不做判定
	_isPassed bool
}

// NewTaobaoqtreportaddRequest 初始化TaobaoqtreportaddAPIRequest对象
func NewTaobaoqtreportaddRequest() *TaobaoqtreportaddAPIRequest {
	return &TaobaoqtreportaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqtreportaddAPIRequest) GetApiMethodName() string {
	return "taobao.qt.report.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqtreportaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqtreportaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServcieItemCode is ServcieItemCode Setter
// 收费项code
func (r *TaobaoqtreportaddAPIRequest) SetServcieItemCode(_servcieItemCode string) error {
	r._servcieItemCode = _servcieItemCode
	r.Set("servcie_item_code", _servcieItemCode)
	return nil
}

// GetServcieItemCode ServcieItemCode Getter
func (r TaobaoqtreportaddAPIRequest) GetServcieItemCode() string {
	return r._servcieItemCode
}

// SetSpName is SpName Setter
// 质检服务商名称
func (r *TaobaoqtreportaddAPIRequest) SetSpName(_spName string) error {
	r._spName = _spName
	r.Set("sp_name", _spName)
	return nil
}

// GetSpName SpName Getter
func (r TaobaoqtreportaddAPIRequest) GetSpName() string {
	return r._spName
}

// SetNick is Nick Setter
// 送检者昵称
func (r *TaobaoqtreportaddAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoqtreportaddAPIRequest) GetNick() string {
	return r._nick
}

// SetQtCode is QtCode Setter
// 一个质检服务唯一标识质量检验单的编号
func (r *TaobaoqtreportaddAPIRequest) SetQtCode(_qtCode string) error {
	r._qtCode = _qtCode
	r.Set("qt_code", _qtCode)
	return nil
}

// GetQtCode QtCode Getter
func (r TaobaoqtreportaddAPIRequest) GetQtCode() string {
	return r._qtCode
}

// SetQtName is QtName Setter
// 质检名称
func (r *TaobaoqtreportaddAPIRequest) SetQtName(_qtName string) error {
	r._qtName = _qtName
	r.Set("qt_name", _qtName)
	return nil
}

// GetQtName QtName Getter
func (r TaobaoqtreportaddAPIRequest) GetQtName() string {
	return r._qtName
}

// SetItemUrl is ItemUrl Setter
// 样品链接.&lt;br/&gt;QT_TYPE=9的时候，请填写N\A
func (r *TaobaoqtreportaddAPIRequest) SetItemUrl(_itemUrl string) error {
	r._itemUrl = _itemUrl
	r.Set("item_url", _itemUrl)
	return nil
}

// GetItemUrl ItemUrl Getter
func (r TaobaoqtreportaddAPIRequest) GetItemUrl() string {
	return r._itemUrl
}

// SetItemDesc is ItemDesc Setter
// 样品信息描述
func (r *TaobaoqtreportaddAPIRequest) SetItemDesc(_itemDesc string) error {
	r._itemDesc = _itemDesc
	r.Set("item_desc", _itemDesc)
	return nil
}

// GetItemDesc ItemDesc Getter
func (r TaobaoqtreportaddAPIRequest) GetItemDesc() string {
	return r._itemDesc
}

// SetQtStandard is QtStandard Setter
// 质检标准&lt;br/&gt;status=3 时 必须非空
func (r *TaobaoqtreportaddAPIRequest) SetQtStandard(_qtStandard string) error {
	r._qtStandard = _qtStandard
	r.Set("qt_standard", _qtStandard)
	return nil
}

// GetQtStandard QtStandard Getter
func (r TaobaoqtreportaddAPIRequest) GetQtStandard() string {
	return r._qtStandard
}

// SetReportUrl is ReportUrl Setter
// 质检报告源文件url&lt;br/&gt;status状态为3时必须非空
func (r *TaobaoqtreportaddAPIRequest) SetReportUrl(_reportUrl string) error {
	r._reportUrl = _reportUrl
	r.Set("report_url", _reportUrl)
	return nil
}

// GetReportUrl ReportUrl Getter
func (r TaobaoqtreportaddAPIRequest) GetReportUrl() string {
	return r._reportUrl
}

// SetMessage is Message Setter
// 检测结果消息描述
func (r *TaobaoqtreportaddAPIRequest) SetMessage(_message string) error {
	r._message = _message
	r.Set("message", _message)
	return nil
}

// GetMessage Message Getter
func (r TaobaoqtreportaddAPIRequest) GetMessage() string {
	return r._message
}

// SetExtAttr is ExtAttr Setter
// 自定义属性字段;分号分隔
func (r *TaobaoqtreportaddAPIRequest) SetExtAttr(_extAttr string) error {
	r._extAttr = _extAttr
	r.Set("ext_attr", _extAttr)
	return nil
}

// GetExtAttr ExtAttr Getter
func (r TaobaoqtreportaddAPIRequest) GetExtAttr() string {
	return r._extAttr
}

// SetGmtSubmit is GmtSubmit Setter
// 送检日期
func (r *TaobaoqtreportaddAPIRequest) SetGmtSubmit(_gmtSubmit string) error {
	r._gmtSubmit = _gmtSubmit
	r.Set("gmt_submit", _gmtSubmit)
	return nil
}

// GetGmtSubmit GmtSubmit Getter
func (r TaobaoqtreportaddAPIRequest) GetGmtSubmit() string {
	return r._gmtSubmit
}

// SetGmtReport is GmtReport Setter
// 提交报告结果时间
func (r *TaobaoqtreportaddAPIRequest) SetGmtReport(_gmtReport string) error {
	r._gmtReport = _gmtReport
	r.Set("gmt_report", _gmtReport)
	return nil
}

// GetGmtReport GmtReport Getter
func (r TaobaoqtreportaddAPIRequest) GetGmtReport() string {
	return r._gmtReport
}

// SetGmtExpiry is GmtExpiry Setter
// 质检有效到期时间，一般为一年有效期&lt;br/&gt;status状态为3时必须非空
func (r *TaobaoqtreportaddAPIRequest) SetGmtExpiry(_gmtExpiry string) error {
	r._gmtExpiry = _gmtExpiry
	r.Set("gmt_expiry", _gmtExpiry)
	return nil
}

// GetGmtExpiry GmtExpiry Getter
func (r TaobaoqtreportaddAPIRequest) GetGmtExpiry() string {
	return r._gmtExpiry
}

// SetQtType is QtType Setter
// (1L, &#34;聚划算&#34;),&lt;br/&gt;(2L, &#34;消保&#34;),&lt;br/&gt;(3L, &#34;分销&#34;),&lt;br/&gt;(4L, &#34;抽检&#34;),&lt;br/&gt;(5L, &#34;良无限线下数据&#34;),&lt;br/&gt;(6L, &#34;入驻/续签商城&#34;),&lt;br/&gt;(7L, &#34;买家质检维权&#34;),&lt;br/&gt;(8L, &#34;实地验证&#34;),&lt;br/&gt;(9L, &#34;淘宝买家订单商品鉴定&#34;),&lt;br/&gt;(10L,&#34;假一赔三&#34;);
func (r *TaobaoqtreportaddAPIRequest) SetQtType(_qtType int64) error {
	r._qtType = _qtType
	r.Set("qt_type", _qtType)
	return nil
}

// GetQtType QtType Getter
func (r TaobaoqtreportaddAPIRequest) GetQtType() int64 {
	return r._qtType
}

// SetStatus is Status Setter
// 0:已提交申请&lt;br/&gt;1:已收到样品&lt;br/&gt;2:已出检测结果&lt;br/&gt;3.已出具报告
func (r *TaobaoqtreportaddAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoqtreportaddAPIRequest) GetStatus() int64 {
	return r._status
}

// SetNumIid is NumIid Setter
// 当前接口只有淘宝订单真假鉴定（QT_TYPE=9）的报告在该字段传入订单号，其他类型报告都不需要传输该值
func (r *TaobaoqtreportaddAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoqtreportaddAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetIsPassed is IsPassed Setter
// 只有status=3时赋值, &lt;br/&gt;true 质检结果合格,false质检结果不合格.&lt;br/&gt;留空表示成分鉴定,不做判定
func (r *TaobaoqtreportaddAPIRequest) SetIsPassed(_isPassed bool) error {
	r._isPassed = _isPassed
	r.Set("is_passed", _isPassed)
	return nil
}

// GetIsPassed IsPassed Getter
func (r TaobaoqtreportaddAPIRequest) GetIsPassed() bool {
	return r._isPassed
}
