package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsAligenieSkillEntityImport 实体动态更新
// alibaba.ailabs.aligenie.skill.entity.import
//
// 根据用户上传的实体信息，进行制定技能实体的动态变更
func AlibabaAilabsAligenieSkillEntityImport(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsAligenieSkillEntityImportAPIRequest, resp *tmallgenie.AlibabaAilabsAligenieSkillEntityImportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
