package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabsaligenieskillentityimport 实体动态更新
// alibaba.ailabs.aligenie.skill.entity.import
//
// 根据用户上传的实体信息，进行制定技能实体的动态变更
func Alibabaailabsaligenieskillentityimport(clt *core.SDKClient, req *tmallgenie.AlibabaailabsaligenieskillentityimportAPIRequest, session string) (*tmallgenie.AlibabaailabsaligenieskillentityimportAPIResponse, error) {
	var resp tmallgenie.AlibabaailabsaligenieskillentityimportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
