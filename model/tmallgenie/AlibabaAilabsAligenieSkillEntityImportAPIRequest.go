package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieSkillEntityImportAPIRequest
实体动态更新 API请求
alibaba.ailabs.aligenie.skill.entity.import

根据用户上传的实体信息，进行制定技能实体的动态变更 */
type AlibabaAilabsAligenieSkillEntityImportAPIRequest struct {
	model.Params
	// 技能Id
	_skillId int64
	// 是否增量更新，true为增量，否则上传数据中没有的实体值将被删除
	_append bool
	// 测试，目前没有此功能，此参数无效
	_test bool
	// 文件内容
	_fileData *model.File
	// 要更新的实体名
	_entityName string
	// 固定值，ISV
	_type string
}

// New
