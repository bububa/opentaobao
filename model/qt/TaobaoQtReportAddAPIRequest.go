package qt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQtReportAddAPIRequest
上传质检报告 API请求
taobao.qt.report.add

上传质检报告 */
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

// New
