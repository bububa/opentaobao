package ihome

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ihome"
)

// AlibabaIhomeCtomCaseMainpicUpdate 方案渲染图修改
// alibaba.ihome.ctom.case.mainpic.update
//
// 用于在门店工作台里的编辑器保存方案，由三维家后端调用阿里后端，保存方案信息
// 此接口只允许ihome业务使用，用于门店的编辑功能，只允许广东三维家信息科技有限公司一家公司调用，不适用于其他业务。
func AlibabaIhomeCtomCaseMainpicUpdate(clt *core.SDKClient, req *ihome.AlibabaIhomeCtomCaseMainpicUpdateAPIRequest, session string) (*ihome.AlibabaIhomeCtomCaseMainpicUpdateAPIResponse, error) {
	var resp ihome.AlibabaIhomeCtomCaseMainpicUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
