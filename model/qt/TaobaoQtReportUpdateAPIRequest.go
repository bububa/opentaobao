package qt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQtReportUpdateAPIRequest 更新质检报告 API请求
// taobao.qt.report.update
//
// 更新质检报告
type TaobaoQtReportUpdateAPIRequest struct {
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
	// 宝贝样品url
	_itemUrl string
	// 样品宝贝描述
	_itemDesc string
	// 质检标准<br/>status=3 时必须非空
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
	// 外部ID，和QT_TYPE 一起表示某种平台的实体ID。QT_TYPE=9的时候，num_iid为淘宝订单号
	_numIid int64
	// 只有status=3时赋值, <br/>0:未通过1:通过 空代表未判定
	_isPassed bool
}

// NewTaobaoQtReportUpdateRequest 初始化TaobaoQtReportUpdateAPIRequest对象
func NewTaobaoQtReportUpdateRequest() *TaobaoQtReportUpdateAPIRequest {
	return &TaobaoQtReportUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQtReportUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.qt.report.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQtReportUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetServcieItemCode is ServcieItemCode Setter
// 收费项code
func (r *TaobaoQtReportUpdateAPIRequest) SetServcieItemCode(_servcieItemCode string) error {
	r._servcieItemCode = _servcieItemCode
	r.Set("servcie_item_code", _servcieItemCode)
	return nil
}

// GetServcieItemCode ServcieItemCode Getter
func (r TaobaoQtReportUpdateAPIRequest) GetServcieItemCode() string {
	return r._servcieItemCode
}

// SetSpName is SpName Setter
// 质检服务商名称
func (r *TaobaoQtReportUpdateAPIRequest) SetSpName(_spName string) error {
	r._spName = _spName
	r.Set("sp_name", _spName)
	return nil
}

// GetSpName SpName Getter
func (r TaobaoQtReportUpdateAPIRequest) GetSpName() string {
	return r._spName
}

// SetNick is Nick Setter
// 送检者昵称
func (r *TaobaoQtReportUpdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoQtReportUpdateAPIRequest) GetNick() string {
	return r._nick
}

// SetQtCode is QtCode Setter
// 一个质检服务唯一标识质量检验单的编号
func (r *TaobaoQtReportUpdateAPIRequest) SetQtCode(_qtCode string) error {
	r._qtCode = _qtCode
	r.Set("qt_code", _qtCode)
	return nil
}

// GetQtCode QtCode Getter
func (r TaobaoQtReportUpdateAPIRequest) GetQtCode() string {
	return r._qtCode
}

// SetQtName is QtName Setter
// 质检名称
func (r *TaobaoQtReportUpdateAPIRequest) SetQtName(_qtName string) error {
	r._qtName = _qtName
	r.Set("qt_name", _qtName)
	return nil
}

// GetQtName QtName Getter
func (r TaobaoQtReportUpdateAPIRequest) GetQtName() string {
	return r._qtName
}

// SetItemUrl is ItemUrl Setter
// 宝贝样品url
func (r *TaobaoQtReportUpdateAPIRequest) SetItemUrl(_itemUrl string) error {
	r._itemUrl = _itemUrl
	r.Set("item_url", _itemUrl)
	return nil
}

// GetItemUrl ItemUrl Getter
func (r TaobaoQtReportUpdateAPIRequest) GetItemUrl() string {
	return r._itemUrl
}

// SetItemDesc is ItemDesc Setter
// 样品宝贝描述
func (r *TaobaoQtReportUpdateAPIRequest) SetItemDesc(_itemDesc string) error {
	r._itemDesc = _itemDesc
	r.Set("item_desc", _itemDesc)
	return nil
}

// GetItemDesc ItemDesc Getter
func (r TaobaoQtReportUpdateAPIRequest) GetItemDesc() string {
	return r._itemDesc
}

// SetQtStandard is QtStandard Setter
// 质检标准&lt;br/&gt;status=3 时必须非空
func (r *TaobaoQtReportUpdateAPIRequest) SetQtStandard(_qtStandard string) error {
	r._qtStandard = _qtStandard
	r.Set("qt_standard", _qtStandard)
	return nil
}

// GetQtStandard QtStandard Getter
func (r TaobaoQtReportUpdateAPIRequest) GetQtStandard() string {
	return r._qtStandard
}

// SetReportUrl is ReportUrl Setter
// 质检报告源文件url&lt;br/&gt;status状态为3时必须非空
func (r *TaobaoQtReportUpdateAPIRequest) SetReportUrl(_reportUrl string) error {
	r._reportUrl = _reportUrl
	r.Set("report_url", _reportUrl)
	return nil
}

// GetReportUrl ReportUrl Getter
func (r TaobaoQtReportUpdateAPIRequest) GetReportUrl() string {
	return r._reportUrl
}

// SetMessage is Message Setter
// 检测结果消息描述
func (r *TaobaoQtReportUpdateAPIRequest) SetMessage(_message string) error {
	r._message = _message
	r.Set("message", _message)
	return nil
}

// GetMessage Message Getter
func (r TaobaoQtReportUpdateAPIRequest) GetMessage() string {
	return r._message
}

// SetExtAttr is ExtAttr Setter
// 自定义属性字段;分号分隔
func (r *TaobaoQtReportUpdateAPIRequest) SetExtAttr(_extAttr string) error {
	r._extAttr = _extAttr
	r.Set("ext_attr", _extAttr)
	return nil
}

// GetExtAttr ExtAttr Getter
func (r TaobaoQtReportUpdateAPIRequest) GetExtAttr() string {
	return r._extAttr
}

// SetGmtSubmit is GmtSubmit Setter
// 送检日期
func (r *TaobaoQtReportUpdateAPIRequest) SetGmtSubmit(_gmtSubmit string) error {
	r._gmtSubmit = _gmtSubmit
	r.Set("gmt_submit", _gmtSubmit)
	return nil
}

// GetGmtSubmit GmtSubmit Getter
func (r TaobaoQtReportUpdateAPIRequest) GetGmtSubmit() string {
	return r._gmtSubmit
}

// SetGmtReport is GmtReport Setter
// 提交报告结果时间
func (r *TaobaoQtReportUpdateAPIRequest) SetGmtReport(_gmtReport string) error {
	r._gmtReport = _gmtReport
	r.Set("gmt_report", _gmtReport)
	return nil
}

// GetGmtReport GmtReport Getter
func (r TaobaoQtReportUpdateAPIRequest) GetGmtReport() string {
	return r._gmtReport
}

// SetGmtExpiry is GmtExpiry Setter
// 质检有效到期时间，一般为一年有效期&lt;br/&gt;status状态为3时必须非空
func (r *TaobaoQtReportUpdateAPIRequest) SetGmtExpiry(_gmtExpiry string) error {
	r._gmtExpiry = _gmtExpiry
	r.Set("gmt_expiry", _gmtExpiry)
	return nil
}

// GetGmtExpiry GmtExpiry Getter
func (r TaobaoQtReportUpdateAPIRequest) GetGmtExpiry() string {
	return r._gmtExpiry
}

// SetQtType is QtType Setter
// (1L, &#34;聚划算&#34;),&lt;br/&gt;(2L, &#34;消保&#34;),&lt;br/&gt;(3L, &#34;分销&#34;),&lt;br/&gt;(4L, &#34;抽检&#34;),&lt;br/&gt;(5L, &#34;良无限线下数据&#34;),&lt;br/&gt;(6L, &#34;入驻/续签商城&#34;),&lt;br/&gt;(7L, &#34;买家质检维权&#34;),&lt;br/&gt;(8L, &#34;实地验证&#34;),&lt;br/&gt;(9L, &#34;淘宝买家订单商品鉴定&#34;),&lt;br/&gt;(10L,&#34;假一赔三&#34;);
func (r *TaobaoQtReportUpdateAPIRequest) SetQtType(_qtType int64) error {
	r._qtType = _qtType
	r.Set("qt_type", _qtType)
	return nil
}

// GetQtType QtType Getter
func (r TaobaoQtReportUpdateAPIRequest) GetQtType() int64 {
	return r._qtType
}

// SetStatus is Status Setter
// 0:已提交申请&lt;br/&gt;1:已收到样品&lt;br/&gt;2:已出检测结果&lt;br/&gt;3.已出具报告
func (r *TaobaoQtReportUpdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoQtReportUpdateAPIRequest) GetStatus() int64 {
	return r._status
}

// SetNumIid is NumIid Setter
// 外部ID，和QT_TYPE 一起表示某种平台的实体ID。QT_TYPE=9的时候，num_iid为淘宝订单号
func (r *TaobaoQtReportUpdateAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoQtReportUpdateAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetIsPassed is IsPassed Setter
// 只有status=3时赋值, &lt;br/&gt;0:未通过1:通过 空代表未判定
func (r *TaobaoQtReportUpdateAPIRequest) SetIsPassed(_isPassed bool) error {
	r._isPassed = _isPassed
	r.Set("is_passed", _isPassed)
	return nil
}

// GetIsPassed IsPassed Getter
func (r TaobaoQtReportUpdateAPIRequest) GetIsPassed() bool {
	return r._isPassed
}
