package ihome

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIhomeCtomCaseMainpicUpdateAPIRequest
方案渲染图修改 API请求
alibaba.ihome.ctom.case.mainpic.update

用于在门店工作台里的编辑器保存方案，由三维家后端调用阿里后端，保存方案信息
此接口只允许ihome业务使用，用于门店的编辑功能，只允许广东三维家信息科技有限公司一家公司调用，不适用于其他业务。 */
type AlibabaIhomeCtomCaseMainpicUpdateAPIRequest struct {
	model.Params
	// 32位字符串
	_traceId string
	// 方案id
	_caseId string
	// 图片的地址
	_picUrl string
	// 图片类型
	_picType string
}

// New
