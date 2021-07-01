package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest
接收中央随机化系统和临床研究机构的绑定确认状态 API请求
alibaba.alihealth.drugcode.center.receive.bound.status

临床用药试验-接收中央随机化系统和临床研究机构的绑定确认状态 */
type AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIRequest struct {
	model.Params
	// 项目id
	_projectId int64
	// 临床研究机构id
	_hospitalRefEntId string
	// 状态 4:绑定成功 5:绑定失败
	_status int64
	// 中央随机化系统id
	_centerRandomSysId string
}

// New
