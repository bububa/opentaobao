package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieOpencontentPushAPIRequest
天猫精灵内容接入标准接口 API请求
alibaba.ailabs.aligenie.opencontent.push

第三方内容接入天猫精灵内容库，供相关技能使用 */
type AlibabaAilabsAligenieOpencontentPushAPIRequest struct {
	model.Params
	// 在Aligenie开放平台创建的技能的ID
	_skillId int64
	// 详细内容列表
	_contents *BatchContent
}

// New
