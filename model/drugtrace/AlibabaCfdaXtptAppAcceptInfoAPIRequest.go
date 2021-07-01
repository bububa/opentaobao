package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCfdaXtptAppAcceptInfoAPIRequest
协同平台数据下行接口 API请求
alibaba.cfda.xtpt.app.accept.info

协同平台数据下行接口 */
type AlibabaCfdaXtptAppAcceptInfoAPIRequest struct {
	model.Params
	// 接入系统唯一标识，由协同平台分配
	_appId string
	// 传输流水号（uuid)
	_processId string
	// 事件编号（uuid）
	_eventId string
	// 事件类型，10：基础信息数据子集 ； 20：应用信息数据子集
	_eventType string
	// 事件子类型
	_subType string
	// 统一社会信用代码
	_uscId string
	// 文件内容 zip压缩+base64转码
	_data string
	// 时间戳，yyyy-MM-dd HH:mm:ss
	_tiemstamp string
}

// New
